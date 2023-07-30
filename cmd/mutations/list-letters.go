package mutations

import (
	"fmt"
	"os"

	"github.com/auenc/geiriadur/mutations"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listLettersCmd = &cobra.Command{
	Use:       "list-letters",
	Aliases:   []string{"lsl"},
	Short:     "lists the letters of the provided mutation",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"all", "soft", "nasal", "aspirate", "h-prothesis"},
	Run: func(cmd *cobra.Command, args []string) {
		mutation := args[0]
		var headers []string

		mutationLetters, err := mutations.MutationLettersAsFlatArray(mutation)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error loading mutation letters '%s'\n", err)
			os.Exit(1)
		}

		headers = []string{"no mutation", mutation}
		if mutation == "all" {
			headers = []string{"no mutation", "soft", "nasal", "aspirate"}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers)
		for _, row := range mutationLetters {
			table.Append(row)
		}

		table.Render()
	},
}
