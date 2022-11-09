package tutorgo

import (
	"fmt"
	"github.com/plavc/tutorial-go-cli/pkg/utillity"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := utillity.Reverse(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
