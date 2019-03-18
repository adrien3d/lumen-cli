package commands

import (
	"github.com/adrien3d/lumen/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func StoreCmd(cmd *cobra.Command, args []string) {
	selectedModels := utils.SelectMethodsModels("store")

	for _, selectedModel := range selectedModels {
		// Step 3.2: Generate interfaces in entity.go
		utils.GenerateFile("entity.store.tmpl", "generated/store/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)

		// Step 3.3: Generate mongo methods in mongodb/entity.go
		utils.GenerateFile("mongo.store.tmpl", "generated/store/mongodb/"+strings.ToLower(selectedModel.ModelName)+".go", selectedModel)
	}

	os.Exit(1)
}
