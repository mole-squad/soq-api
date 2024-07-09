package main

import (
	"context"

	"github.com/burkel24/task-app/pkg/app"
	"github.com/burkel24/task-app/pkg/interfaces"
	"go.uber.org/fx"
)

func GenerateAgendas(lc fx.Lifecycle, agendaService interfaces.AgendaService) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := agendaService.GenerateAgendasForUpcomingTimeWindows(ctx)
	if err != nil {
		panic(err)
	}

	err = agendaService.PopulatePendingAgendas(ctx)
	if err != nil {
		panic(err)
	}
}

func main() {
	appOpts := app.BuildAppOpts()
	appOpts = append(appOpts, fx.Invoke(GenerateAgendas))

	// TODO figure out how to stop it once GenerateAgendas is done
	fx.New(appOpts...).Run()
}
