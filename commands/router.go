package commands

import (
	"bytes"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/adrien3d/lumen/utils"
	"github.com/gedex/inflector"
	"github.com/spf13/cobra"
	"go/format"
	"gopkg.in/godo.v2/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func generateRouterFile(selectedModels []SelectedModel) {
	for _, model := range selectedModels {
		fmt.Print(model.ModelName, ": ")
		for _, method := range model.Methods {
			fmt.Print(method, "\t")
		}
		fmt.Println()
	}

	path := filepath.Join("templates", "router.tmpl")
	body, _ := ioutil.ReadFile(path)
	tmpl := template.Must(template.New("model").Option("missingkey=error").Funcs(funcMap).Parse(string(body)))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, selectedModels)
	utils.Check(err)

	src, _ := format.Source(buf.Bytes())
	dstPath := filepath.Join("generated/server/router.go")

	if !util.FileExists(filepath.Dir(dstPath)) {
		if err := os.Mkdir(filepath.Dir(dstPath), 0644); err != nil {
			fmt.Println(err)
		}
	}
	if err := ioutil.WriteFile(dstPath, src, 0644); err != nil {
		fmt.Println(err)
	}
}

func RouterCmd(cmd *cobra.Command, args []string) {
	shell := ishell.New()

	shell.Println("Generating Router")

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
		"Please select models you want to generate the matching router:",
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

	generateRouterFile(selectedModels)

	os.Exit(1)

	shell.Run()
}
