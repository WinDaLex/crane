package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export manifests to the current working directory",
	Run: func(cmd *cobra.Command, args []string) {
		for _, manifest := range args {
			if err := manager.Export(manifest); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
}
