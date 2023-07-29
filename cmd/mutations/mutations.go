package mutations

import (
	"github.com/spf13/cobra"
)

var MutationsCmd = &cobra.Command{
	Use:     "mutations",
	Aliases: []string{"m"},
	Short:   "A set of commands related to mutations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	MutationsCmd.AddCommand(mutateCmd)
	MutationsCmd.AddCommand(listLettersCmd)
}
