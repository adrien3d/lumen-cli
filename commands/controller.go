package commands

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func generateControllerFile(methods []string) {
	/*methodSlice := make([]string, len(methods))
	for _, method := range methods {
		methodSlice = append(methodSlice, method)
	}*/

	for _, method := range methods {
		switch method {
		case "Create":
			break
		case "Get":
			break
		case "Update":
			break
		case "Delete":
			break
		}
	}
}

func ControllerCmd(cmd *cobra.Command, args []string) {
	//fmt.Println("Generating Controller: " + strings.Join(args, " "))
	shell := ishell.New()

	shell.Println("Generating Controller")

	files, err := ioutil.ReadDir("generated/models")
	if err != nil {
		fmt.Println(err)
	}

	var fileNames []string
	for _, file := range files {
		fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name())
	}

	choice := shell.MultiChoice(fileNames, "Please select ?")

	shell.Println(fileNames[choice])

	methods := []string{"Create", "Get", "Update", "Delete"}
	choices := shell.Checklist(methods,
		"What method do you want to implement ? (space to select/deselect)",
		nil)
	out := func() (c []string) {
		for _, v := range choices {
			c = append(c, methods[v])
		}
		return
	}
	selected := out()
	//shell.Println("Your choices are", strings.Join(out(), ", "))

	/*methodSlice := make([]string, len(selected))
	for _, method := range selected {
		methodSlice = append(methodSlice, method)
	}*/

	generateControllerFile( /*methodSlice*/ selected)

	shell.Run()
}
