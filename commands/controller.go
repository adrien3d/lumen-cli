package commands

import (
	"github.com/adrien3d/lumen/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// ControllerCmd holds functions to generate controller
func ControllerCmd(cmd *cobra.Command, args []string) {
	selectedModels := utils.SelectMethodsModels("controller")

	for _, selectedModel := range selectedModels {
		utils.GenerateFile("controller.tmpl", "generated/controllers/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)
	}

	os.Exit(1)
}
