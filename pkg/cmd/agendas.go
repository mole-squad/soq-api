package cmd

import (
	"context"
	"os"

	"github.com/burkel24/task-app/pkg/app"
	"github.com/burkel24/task-app/pkg/interfaces"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var agendasCmd = &cobra.Command{
	Use: "agendas",
	Run: func(cmd *cobra.Command, args []string) {
		appOpts := app.BuildAppOpts()
		appOpts = append(appOpts, fx.Invoke(GenerateAndSendAgendas))

		fx.New(appOpts...).Run()
	},
}

type GenerateAgendasParams struct {
	fx.In

	AgendaService interfaces.AgendaService
	Logger        interfaces.LoggerService
}

func GenerateAndSendAgendas(params GenerateAgendasParams) error {
	params.Logger.Info("Generating agendas")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := params.AgendaService.GenerateAgendasForUpcomingTimeWindows(ctx)
	if err != nil {
		return err
	}

	err = params.AgendaService.PopulatePendingAgendas(ctx)
	if err != nil {
		return err
	}

	params.Logger.Info("Finished generating agendas")
	params.Logger.Info("Sending agenda notifications")

	err = params.AgendaService.SendAgendaNotifications(ctx)
	if err != nil {
		return err
	}

	params.Logger.Info("Finished sending agenda notifications")

	os.Exit(0)

	return nil
}
