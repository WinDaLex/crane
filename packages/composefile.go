package packages

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type composefile string

func (c composefile) validate() ([]string, []string, error) {
	out, err := exec.Command("docker-compose", "-f", string(c), "config").CombinedOutput()
	if err != nil {
		return nil, nil, errors.New("composefile format error")
	}

	config := string(out)

	imagesMap := map[string]bool{}
	for _, line := range strings.Split(config, "\n") {
		if strings.Contains(line, "image: ") {
			image := strings.TrimPrefix(strings.TrimSpace(line), "image: ")
			imagesMap[image] = true
		}
	}

	images := []string{}
	for i := range imagesMap {
		images = append(images, i)
	}

	// A service in valid Docker Compose file should have either an image or a build context specified. If no image is provided, there should be a build context in the Compose file, that is not allowed in a Crane manifest.
	if len(images) == 0 {
		return nil, nil, errors.New("should specfic at least an image")
	}

	out, _ = exec.Command("docker-compose", "-f", string(c), "config", "--services").Output()

	commands := strings.Split(string(out), "\n")
	commands = commands[:len(commands)-1]

	return commands, images, nil
}
