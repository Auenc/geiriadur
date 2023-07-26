package geiriadur

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geiriadur",
	Short: "geiriadur - a simple welsh dictionary cli that will detail different forms of a word",
	Long:  "geiriadur - a simple welsh dictionary cli that will detail different forms of a word",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("S'mae, byd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
