package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func RouterCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Generating Router: " + strings.Join(args, " "))
}
