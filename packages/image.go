package packages

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/windalex/crane/path"
)

func init() {
	os.Setenv("DOCKER_HOST", "unix://"+path.DockerHost())
}

type img string

func (i img) pull() error {
	cmd := exec.Command("docker", "pull", string(i))

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "docker pull %s", string(i))
	}
	return nil
}

func (i img) remove() error {
	cmd := exec.Command("docker", "rmi", "-f", string(i))

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "docker rmi -f %s", string(i))
	}
	return nil
}
