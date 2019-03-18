package commands

import (
	"github.com/adrien3d/lumen/utils"
	"github.com/spf13/cobra"
	"os"
)

func RouterCmd(cmd *cobra.Command, args []string) {
	selectedModels := utils.SelectMethodsModels("router")

	utils.GenerateFile("router.tmpl", "generated/server/router.go", selectedModels)

	os.Exit(1)
}
