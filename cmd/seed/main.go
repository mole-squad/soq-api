package main

import (
	"context"

	"github.com/burkel24/task-app/pkg/app"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/burkel24/task-app/pkg/models"
	"go.uber.org/fx"
)

func NewSeeder(lc fx.Lifecycle, dbService interfaces.DBService) {
	dbService.CreateOne(context.Background(), &models.User{})
}

func main() {
	serverOpts := app.BuildAppOpts()
	serverOpts = append(serverOpts, fx.Invoke(NewSeeder))

	fx.New(serverOpts...).Run()
}
