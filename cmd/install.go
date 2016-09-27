package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var (
	prefix string
	suffix string
)

var installCmd = &cobra.Command{
	Use:   "install PACKAGE [PACKAGE...]",
	Short: "Install new packages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "crane: \"install\" requires at least one argument.")
			fmt.Fprintln(os.Stderr, "See 'crane install --help'")
		} else {
			for _, pkg := range args {
				if err := manager.Install(pkg, "", prefix, suffix); err != nil {
					fmt.Fprintln(os.Stderr, errors.Wrapf(err, "install %s failed", pkg))
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "Prefix")
	installCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "Suffix")
}
