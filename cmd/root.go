package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd root comand for cobra
var RootCmd = &cobra.Command{
	Use:           "ecs-goploy",
	Short:         "Deploy commands for ECS",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		serviceCmd(),
		taskCmd(),
		versionCmd(),
		taskDefinitionCmd(),
	)
}
