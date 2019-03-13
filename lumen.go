package main

import (
	"github.com/adrien3d/lumen/commands"
	"github.com/spf13/cobra"
)

func main() {
	var cmdController = &cobra.Command{
		Use:   "controller [name]",
		Short: `Generating controller`,
		Long:  `Generates a controller with multiple methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.ControllerCmd,
	}

	var cmdModel = &cobra.Command{
		Use:   "model [name]",
		Short: `Generating model`,
		Long:  "Generates a model",
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.ModelCmd,
	}

	var cmdRouter = &cobra.Command{
		Use:   "router [name]",
		Short: `Generating router`,
		Long:  `Generates router methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.RouterCmd,
	}

	var cmdStore = &cobra.Command{
		Use:   "store [name]",
		Short: `Generating store`,
		Long:  `Generates store methods`,
		Args:  cobra.MinimumNArgs(0),
		Run:   commands.StoreCmd,
	}

	var rootCmd = &cobra.Command{
		Use:   "lumen",
		Short: "Lumen is a CLI that helps you generating API code for base-api"}
	rootCmd.AddCommand(cmdController, cmdModel, cmdRouter, cmdStore)
	rootCmd.Execute()
}
