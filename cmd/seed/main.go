package main

import (
	"context"
	"time"

	"github.com/burkel24/task-app/pkg/app"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
)

var (
	weekdays = []int32{int32(time.Monday), int32(time.Tuesday), int32(time.Wednesday), int32(time.Thursday), int32(time.Friday)}
	weekends = []int32{int32(time.Saturday), int32(time.Sunday)}
)

type SeederParams struct {
	fx.In

	DbService interfaces.DBService
	Logger    interfaces.LoggerService
}

func NewSeeder(params SeederParams) {
	params.Logger.Info("Seeding database")

	params.DbService.DropAll(context.Background())
	params.DbService.Migrate(context.Background())

	user := models.User{Name: "Burke", Timezone: "America/Los_Angeles"}
	params.DbService.CreateOne(context.Background(), &user)

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
		Summary:     "Write Tests",
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
}

func main() {
	appOpts := app.BuildAppOpts()
	appOpts = append(appOpts, fx.Invoke(NewSeeder))

	fx.New(appOpts...).Run()
}
