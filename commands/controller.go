package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func GenerateController(cmd *cobra.Command, args []string) {
	fmt.Println("Generating Controller: " + strings.Join(args, " "))
}