package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/windalex/crane/manager"
)

var (
	avaiable bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed packages",
	Run: func(cmd *cobra.Command, args []string) {
		var info []string
		var err error

		if avaiable {
			info, err = manager.AllManifests()
		} else {
			info, err = manager.AllInstalledPackages()
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "failed")
		} else {
			fmt.Println(strings.Join(info, "\n"))
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&avaiable, "avaiable", false, "Help message for toggle")
}
