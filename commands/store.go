package commands

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/adrien3d/lumen/utils"
	"github.com/gedex/inflector"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

func StoreCmd(cmd *cobra.Command, args []string) {
	shell := ishell.New()

	shell.Println("Generating Store")

	files, err := ioutil.ReadDir("generated/models")
	if err != nil {
		fmt.Println(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, strings.Title(inflector.Singularize(strings.TrimRight(file.Name(), ".go"))))
	}

	// Step 1: Ask user which models to use
	choices := shell.Checklist(fileNames,
		"Please select models you want to generate the matching store:",
		nil)
	fmt.Println(choices)

	var selectedModels []SelectedModel
	for _, file := range choices {
		var selectedModel SelectedModel
		selectedModel.ModelName = fileNames[file]
		//Step 2: Select methods to generate
		methods := []string{"Create" + fileNames[file], "Get" + fileNames[file], "GetAll" + fileNames[file], "Update" + fileNames[file], "Delete" + fileNames[file]}
		choices := shell.Checklist(methods,
			"What method do you want to implement ? (space to select/deselect)",
			nil)
		for _, v := range choices {
			meth := methods[v]
			meth = meth[0 : len(meth)-len(fileNames[file])]
			selectedModel.Methods = append(selectedModel.Methods, meth)
		}
		selectedModels = append(selectedModels, selectedModel)
	}

	// Step 1: Generate store.go that indexes methods
	utils.GenerateFile("index.store.tmpl", "generated/store/store.go", selectedModels)

	// Step 2: Generate interfaces in entity.go

	os.Exit(1)

	shell.Run()
}
