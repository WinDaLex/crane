package packages

import (
	"os/exec"

	"github.com/pkg/errors"
	"github.com/windalex/crane/path"
)

func Uninstall(name string) error {
	p, err := retrive(name)
	if err != nil {
		return errors.Errorf("package is not installed")
	}

	for _, c := range p.commands {
		cmd(p.prefix + c + p.suffix).destroy()
	}

	exec.Command("docker-compose", "-f", path.PackageComposefile(p.name), "down", "-v").Run()

	for _, i := range p.images {
		img(i).remove()
	}

	p.destroy()

	return nil
}
