package mutations

import (
	"os"

	"github.com/auenc/geiriadur/mutations"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var mutateCmd = &cobra.Command{
	Use:     "mutate",
	Aliases: []string{"m"},
	Short:   "list the mutations for a provided word",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]
		mutation, err := mutations.Mutate(word)
		if err != nil {
			panic(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Mutation", "Word"})

		table.Append([]string{"-", word})
		table.Append([]string{"Soft", string(mutation.Soft)})
		table.Append([]string{"Nasal", string(mutation.Nasal)})
		table.Append([]string{"Aspirate", string(mutation.Aspirate)})
		table.Append([]string{"H-Prothesis", string(mutation.HProthesis)})

		table.Render()
	},
}
