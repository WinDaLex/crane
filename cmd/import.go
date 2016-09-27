package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import manifests",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, manifest := range args {
				fmt.Fprintln(os.Stderr, manager.Import(manifest))
			}
		} else {
			fmt.Fprintln(os.Stderr, manager.Import("docker-compose.yml"))
		}
	},
}

func init() {
	RootCmd.AddCommand(importCmd)
}
