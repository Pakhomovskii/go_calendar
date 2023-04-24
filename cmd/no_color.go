package cmd

import (
	"github.com/spf13/cobra"
)

var noColorCmd = &cobra.Command{
	Use:   "no-color",
	Short: "Disables colorized output",
	Run: func(cmd *cobra.Command, args []string) {
		generateOutput(true)
	},
}

func init() {
	rootCmd.AddCommand(noColorCmd)
}
