package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall packages",
	Run: func(cmd *cobra.Command, args []string) {
		for _, pkg := range args {
			if err := manager.Uninstall(pkg); err != nil {
				fmt.Fprintf(os.Stderr, "uninstall %s failed: %v\n", pkg, err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(uninstallCmd)
}
