package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/abiosoft/ishell"
	"github.com/gedex/inflector"
	"github.com/serenize/snaker"
	"gopkg.in/godo.v2/util"
)

// Check checks errors
func Check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// FirstCharLower lowers first char of string
func FirstCharLower(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

// GetFirstChar returns the first char of string
func GetFirstChar(s string) string {
	return s[0:1]
}

// FuncMap is a set of functions to use in templates
var FuncMap = template.FuncMap{
	"pluralize":   inflector.Pluralize,
	"singularize": inflector.Singularize,
	"title":       strings.Title,
	"firstLower":  FirstCharLower,
	"toLower":     strings.ToLower,
	"toSnakeCase": snaker.CamelToSnake,
	"firstChar":   GetFirstChar,
}

// SelectedModel is a structure that holds model name and selected methods
type SelectedModel struct {
	Namespace string
	ModelName string
	Methods   []string
}

// SelectMethodsModels asks user which models and methods to generate
func SelectMethodsModels(typeName string) (selectedModels []SelectedModel) {
	shell := ishell.New()

	shell.Println("Generating " + typeName)

	shell.Print("Namespace: ")
	namespace := shell.ReadLine()

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

	for _, file := range choices {
		var selectedModel SelectedModel
		selectedModel.Namespace = namespace
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

	return selectedModels
}

// GenerateFile produces a file from template and data (data structure)
func GenerateFile(templateFile string, outputPath string, data interface{}) {
	path := filepath.Join("templates", templateFile)
	body, _ := ioutil.ReadFile(path)
	tmpl := template.Must(template.New("model").Option("missingkey=error").Funcs(FuncMap).Parse(string(body)))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	Check(err)

	src, _ := format.Source(buf.Bytes())
	dstPath := filepath.Join(outputPath)

	if !util.FileExists(filepath.Dir(dstPath)) {
		if err := os.Mkdir(filepath.Dir(dstPath), 0644); err != nil {
			fmt.Println(err)
		}
	}
	if err := ioutil.WriteFile(dstPath, src, 0644); err != nil {
		fmt.Println(err)
	}
}
