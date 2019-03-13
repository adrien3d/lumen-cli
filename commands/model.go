package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func GenerateModel(cmd *cobra.Command, args []string) {
	fmt.Println("Generating Model: " + strings.Join(args, " "))
}
