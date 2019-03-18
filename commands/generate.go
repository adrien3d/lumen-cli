package commands

import (
	"github.com/adrien3d/lumen/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// GenerateCmd holds functions to generate controllers, router, store
func GenerateCmd(cmd *cobra.Command, args []string) {
	selectedModels := utils.SelectMethodsModels("everything")

	// Step 1: Generating controllers
	for _, selectedModel := range selectedModels {
		utils.GenerateFile("controller.tmpl", "generated/controllers/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)
	}

	// Step 2: Generating router
	utils.GenerateFile("router.tmpl", "generated/server/router.go", selectedModels)

	// Step 3.1: Generate store.go that indexes methods
	utils.GenerateFile("index.store.tmpl", "generated/store/store.go", selectedModels)

	for _, selectedModel := range selectedModels {
		// Step 3.2: Generate interfaces in entity.go
		utils.GenerateFile("entity.store.tmpl", "generated/store/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)

		// Step 3.3: Generate mongo methods in mongodb/entity.go
		utils.GenerateFile("mongo.store.tmpl", "generated/store/mongodb/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)
	}

	os.Exit(1)

}
