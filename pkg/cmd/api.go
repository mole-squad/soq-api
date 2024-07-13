package cmd

import (
	"github.com/mole-squad/soq/pkg/app"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var apiCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		appOpts := app.BuildAppOpts()
		serverOpts := app.BuildServerOpts()

		allOpts := append(appOpts, serverOpts...)

		fx.New(allOpts...).Run()
	},
}
