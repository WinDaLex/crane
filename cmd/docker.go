package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv("DOCKER_HOST", "unix:///home/vagrant/.crane/var/run/docker.sock")

		command := exec.Command("docker", args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to run docker command: %v\n", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(dockerCmd)
}
