package main

import (
	"github.com/burkel24/task-app/pkg/app"
	"go.uber.org/fx"
)

func main() {
	appOpts := app.BuildAppOpts()
	serverOpts := app.BuildServerOpts()

	allOpts := append(appOpts, serverOpts...)

	fx.New(allOpts...).Run()
}
