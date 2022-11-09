package tutorgo

import (
	"fmt"
	"github.com/plavc/tutorial-go-cli/pkg/hash"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Generate hash",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := hash.Sha256(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
}
