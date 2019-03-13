package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func StoreCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Generating Store: " + strings.Join(args, " "))
}
