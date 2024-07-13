package cmd

import (
	"context"
	"os"
	"time"

	"github.com/mole-squad/soq/pkg/app"
	"github.com/mole-squad/soq/pkg/interfaces"
	"github.com/mole-squad/soq/pkg/models"
	"github.com/mole-squad/soq/pkg/users"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var seedCmd = &cobra.Command{
	Use: "seed",
	Run: func(cmd *cobra.Command, args []string) {
		appOpts := app.BuildAppOpts()
		appOpts = append(appOpts, fx.Invoke(NewSeeder))

		fx.New(appOpts...).Run()
	},
}

var (
	weekdays = []int32{int32(time.Monday), int32(time.Tuesday), int32(time.Wednesday), int32(time.Thursday), int32(time.Friday)}
	weekends = []int32{int32(time.Saturday), int32(time.Sunday)}
)

type SeederParams struct {
	fx.In

	DbService interfaces.DBService
	Logger    interfaces.LoggerService
	interfaces.UserService
}

func NewSeeder(params SeederParams) error {
	params.Logger.Info("Seeding database")

	params.DbService.DropAll(context.Background())
	params.DbService.Migrate(context.Background())

	password := os.Getenv("TEST_USER_PASSWORD")
	hash, err := users.HashUserPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username:     "burke",
		PasswordHash: hash,
		Name:         "Burke",
		Timezone:     "America/Los_Angeles",
	}

	params.DbService.CreateOne(context.Background(), &user)

	params.DbService.CreateOne(context.Background(), &models.Device{
		UserKey:  os.Getenv("PUSHOVER_USER_KEY"),
		UserID:   user.ID,
		DeviceID: "burke-iphone",
	})

	workFocusArea := models.FocusArea{
		Name:   "Work",
		UserID: user.ID,
	}

	personalFocusArea := models.FocusArea{
		Name:   "Personal",
		UserID: user.ID,
	}

	params.DbService.CreateOne(context.Background(), &workFocusArea)
	params.DbService.CreateOne(context.Background(), &personalFocusArea)

	params.DbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   9.0,
		EndTime:     17.0,
		FocusAreaID: workFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   6.0,
		EndTime:     9.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   12.0,
		EndTime:     13.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   18.0,
		EndTime:     22.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekends,
		StartTime:   8.0,
		EndTime:     20.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.Task{
		Summary:     "Write Code",
		Status:      models.TaskStatusOpen,
		FocusAreaID: workFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.Task{
		Summary:     "Exercise",
		Status:      models.TaskStatusOpen,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	params.DbService.CreateOne(context.Background(), &models.Task{
		Summary:     "Do Laundry",
		Status:      models.TaskStatusOpen,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	quota := models.Quota{
		Summary:     "Work Out",
		Period:      models.DailyQuota,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	}

	params.DbService.CreateOne(context.Background(), &quota)

	params.Logger.Info("Database seeded")

	os.Exit(0)
	return nil
}
