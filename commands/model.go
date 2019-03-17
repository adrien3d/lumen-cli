package commands

import (
	"bytes"
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/adrien3d/lumen/utils"
	"github.com/gedex/inflector"
	"github.com/serenize/snaker"
	"github.com/spf13/cobra"
	"go/format"
	"gopkg.in/godo.v2/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Model struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name     string
	Type     string
	IsId     bool
	Required bool
}

var funcMap = template.FuncMap{
	"pluralize":        inflector.Pluralize,
	"title":            strings.Title,
	"toLower":          strings.ToLower,
	"toLowerCamelCase": utils.CamelToLowerCamel,
	"toOriginalCase":   utils.CamelToOriginal,
	"toSnakeCase":      snaker.CamelToSnake,
	"firstChar":        utils.GetFirstChar,
}

func generateModelFile(model Model) {
	path := filepath.Join("templates", "model.tmpl")
	body, _ := ioutil.ReadFile(path)
	tmpl, _ := template.New("model").Funcs(funcMap).Parse(string(body))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, model)
	utils.Check(err)

	src, _ := format.Source(buf.Bytes())
	dstPath := filepath.Join("generated/models/", snaker.CamelToSnake(model.Name)+".go")

	if !util.FileExists(filepath.Dir(dstPath)) {
		if err := os.Mkdir(filepath.Dir(dstPath), 0644); err != nil {
			fmt.Println(err)
		}
	}

	if err := ioutil.WriteFile(dstPath, src, 0644); err != nil {
		fmt.Println(err)
	}
}

func ModelCmd(cmd *cobra.Command, args []string) {
	var fields []*Field
	model := Model{"", fields}

	//fmt.Println("Generating Model: " + strings.Join(args, " "))
	shell := ishell.New()

	// Step 0: choosing model name, if not already done
	if len(args) == 1 {
		shell.Print("Generating Model ")
		shell.Println(`"` + args[0] + `"`)
		model.Name = args[0]
	} else {
		shell.Println("Generating Model")
		shell.Print("Name: ")
		model.Name = shell.ReadLine()
	}

	propertyNum := 0
	for propertyNum < 50 {
		// Step 1: choosing property name
		shell.Print("Enter the property name: ")
		field := Field{"", "", false, false}
		field.Name = shell.ReadLine()

		// Step 2: choosing property type
		choice := shell.MultiChoice([]string{
			"string",
			"int",
			"float",
			"bool",
		}, "Property "+field.Name+" Type: ")

		switch choice {
		case 0:
			field.Type = "string"
			break
		case 1:
			field.Type = "int"
			break
		case 2:
			field.Type = "float"
			break
		case 3:
			field.Type = "bool"
			break
		default:
			field.Type = "string"
			break
		}
		shell.Println(field.Type)

		// Step 3: choosing if property is the ID one
		if shell.MultiChoice([]string{"yes", "no"}, "Is "+field.Name+" the ID property?") == 0 {
			field.IsId = true
		}

		// Step 4: choosing if property is required
		if shell.MultiChoice([]string{"yes", "no"}, "Is "+field.Name+" required?") == 0 {
			field.Required = true
		}

		model.Fields = append(model.Fields, &field)

		propertyNum += 1

		if shell.MultiChoice([]string{"yes", "no"}, "Do you want to add a new property?") == 1 {
			generateModelFile(model)
			os.Exit(1)
		}
	}

	shell.Run()
}
