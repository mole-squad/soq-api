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

func NewSeeder(lc fx.Lifecycle, dbService interfaces.DBService) {
	dbService.DropAll(context.Background())
	dbService.Migrate(context.Background())

	user := models.User{Name: "Burke", Timezone: "America/Los_Angeles"}
	dbService.CreateOne(context.Background(), &user)

	workFocusArea := models.FocusArea{
		Name:   "Work",
		UserID: user.ID,
	}

	personalFocusArea := models.FocusArea{
		Name:   "Personal",
		UserID: user.ID,
	}

	dbService.CreateOne(context.Background(), &workFocusArea)
	dbService.CreateOne(context.Background(), &personalFocusArea)

	dbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   9.0,
		EndTime:     17.0,
		FocusAreaID: workFocusArea.ID,
		UserID:      user.ID,
	})

	dbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   6.0,
		EndTime:     9.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	dbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   12.0,
		EndTime:     13.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	dbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekdays,
		StartTime:   18.0,
		EndTime:     22.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	dbService.CreateOne(context.Background(), &models.TimeWindow{
		Weekdays:    weekends,
		StartTime:   8.0,
		EndTime:     20.0,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	})

	quota := models.Quota{
		Summary:     "Work Out",
		Period:      models.DailyQuota,
		FocusAreaID: personalFocusArea.ID,
		UserID:      user.ID,
	}

	dbService.CreateOne(context.Background(), &quota)
}

func main() {
	serverOpts := app.BuildAppOpts()
	serverOpts = append(serverOpts, fx.Invoke(NewSeeder))

	fx.New(serverOpts...).Run()
}
