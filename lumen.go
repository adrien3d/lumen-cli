package main

import (
	"github.com/adrien3d/lumen/commands"
	"github.com/spf13/cobra"
)

func main() {
	var cmdController = &cobra.Command{
		Use:   "controller [name]",
		Short: `Generating controller`,
		Long: `Generates a controller with multiple methods`,
		Args: cobra.MinimumNArgs(1),
		Run: commands.GenerateController,
	}

	var cmdModel = &cobra.Command{
		Use:   "model [name]",
		Short: `Generating model`,
		Long: "Generates a model",
		Args: cobra.MinimumNArgs(1),
		Run: commands.GenerateModel,
	}

	var cmdRouter = &cobra.Command{
		Use:   "router [name]",
		Short: `Generating router`,
		Long: `Generates router methods`,
		Args: cobra.MinimumNArgs(1),
		Run: commands.GenerateRouter,
	}

	var cmdStore = &cobra.Command{
		Use:   "store [name]",
		Short: `Generating store`,
		Long: `Generates store methods`,
		Args: cobra.MinimumNArgs(1),
		Run: commands.GenerateStore,
	}

	var rootCmd = &cobra.Command{Use: "lumen"}
	rootCmd.AddCommand(cmdController, cmdModel, cmdRouter, cmdStore)
	rootCmd.Execute()
}
