package cmd

import (
	"fmt"
	"os"

	"github.com/auenc/geiriadur/cmd/mutations"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "geiriadur",
	Short: "geiriadur - a simple welsh dictionary cli that will detail different forms of a word",
	Long:  "geiriadur - a simple welsh dictionary cli that will detail different forms of a word",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(mutations.MutationsCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
