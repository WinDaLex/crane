package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var (
	clear bool
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show the history of package management operations",
	Run: func(cmd *cobra.Command, args []string) {
		if clear {
			manager.ClearHistory()
		} else {
			fmt.Fprintln(os.Stderr, manager.ShowHistory())
		}
	},
}

func init() {
	RootCmd.AddCommand(historyCmd)

	historyCmd.Flags().BoolVar(&clear, "clear", false, "Clear")
}
