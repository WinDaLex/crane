package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade packages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, arg := range args {
				if err := manager.UpgradePackage(arg); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		} else {
			if err := manager.UpgradeAllPackages(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(upgradeCmd)
}
