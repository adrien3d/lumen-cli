package commands

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func GenerateController(cmd *cobra.Command, args []string) {
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

	shell.Run()
}
