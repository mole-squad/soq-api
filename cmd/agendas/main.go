package main

import (
	"context"

	"github.com/burkel24/task-app/pkg/app"
	"github.com/burkel24/task-app/pkg/interfaces"
	"go.uber.org/fx"
)

type GenerateAgendasParams struct {
	fx.In

	AgendaService interfaces.AgendaService
	Logger        interfaces.LoggerService
}

func GenerateAgendas(params GenerateAgendasParams) {
	params.Logger.Info("Generating agendas")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := params.AgendaService.GenerateAgendasForUpcomingTimeWindows(ctx)
	if err != nil {
		panic(err)
	}

	err = params.AgendaService.PopulatePendingAgendas(ctx)
	if err != nil {
		panic(err)
	}

	params.Logger.Info("Finished generating agendas")
}

func main() {
	appOpts := app.BuildAppOpts()
	appOpts = append(appOpts, fx.Invoke(GenerateAgendas))

	// TODO figure out how to stop it once GenerateAgendas is done
	fx.New(appOpts...).Run()
}
