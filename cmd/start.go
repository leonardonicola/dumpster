/*
Copyright © 2024 Leonardo Nicola
*/
package cmd

import (
	"github.com/leonardonicola/dumpster/internal"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the backups right now!",
	Long: `You can choose between our databases options and start a interactive form, providing the credentials for the database in order to start the backups. 

With dumpster you make multiple backups at once in parallel, which allows us to extract the most performance out the process.`,

	Run: func(cmd *cobra.Command, args []string) {
		internal.HandleDatabaseSelection()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
