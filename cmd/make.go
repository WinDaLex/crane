package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var (
	makePrefix string
	makeSuffix string
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Install new packages by giving manifests",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, manifest := range args {
				if err := manager.Install("", manifest, makePrefix, makeSuffix); err != nil {
					fmt.Fprintln(os.Stderr, errors.Wrapf(err, "install from %s failed", manifest))
				}
			}
		} else {
			if err := manager.Install("", "docker-compose.yml", makePrefix, makeSuffix); err != nil {
				fmt.Fprintln(os.Stderr, errors.Wrapf(err, "install from %s failed", "docker-compose.yml"))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(makeCmd)

	makeCmd.Flags().StringVarP(&makePrefix, "prefix", "p", "", "Prefix")
	makeCmd.Flags().StringVarP(&makeSuffix, "suffix", "s", "", "Suffix")
}
