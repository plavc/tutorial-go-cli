package tutorgo

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const Version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "tutorgo",
	Version: Version,
	Short:   "tutorgo - a simple CLI template implemented in Go",
	Long: `
'tutorgo' is a simple CLI template implemented in Go

One can use 'tutorgo' to implement advanced cli tool`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing tutorgo '%s'", err)
		os.Exit(1)
	}
}
