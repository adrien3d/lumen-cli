package commands

import (
	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"
)

const (
	AssociationNone      = 0
	AssociationBelongsTo = 1
	AssociationHasMany   = 2
	AssociationHasOne    = 3
)

type Model struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name        string
	JSONName    string
	Type        string
	Tag         string
	Association *Association
}

type Association struct {
	Type  int
	Model *Model
}

func GenerateModel(cmd *cobra.Command, args []string) {
	fields := []*Field{}
	model := Model{"", fields}

	//fmt.Println("Generating Model: " + strings.Join(args, " "))
	shell := ishell.New()

	if len(args) == 1 {
		shell.Print("Generating Controller ")
		shell.Println(`"` + args[0] + `"`)
		model.Name = args[0]
	} else {
		shell.Println("Generating Controller")
		shell.Print("Name: ")
		model.Name = shell.ReadLine()
	}

	choice := shell.MultiChoice([]string{
		"string",
		"int",
		"bool",
	}, "Type ?")

	switch choice {
	case 0:
		shell.Println("string")
		break
	case 1:
		shell.Println("int")
		break
	case 2:
		shell.Println("bool")
		break
	default:
		shell.Println("string")
		break
	}

	shell.Run()
}
