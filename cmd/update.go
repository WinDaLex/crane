package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update manifests",
	Run: func(cmd *cobra.Command, args []string) {
		if err := manager.UpdateManifests(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
