package packages

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/windalex/crane/path"
)

func Install(name string, manifest string, prefix string, suffix string) error {
	cmds, imgs, err := composefile(manifest).validate()
	if err != nil {
		return errors.Wrapf(err, "validate %s failed", manifest)
	}

	for _, c := range cmds {
		if cmd(c).exists() {
			return errors.Errorf("command %s conflicts", c)
		}
	}

	fmt.Printf("installing '%s' ...\n", name)

	for _, i := range imgs {
		fmt.Printf("Downloading dependent image '%s' ...", i)
		if err := img(i).pull(); err != nil {
			return errors.Wrapf(err, "pull image %s failed", i)
		}
		fmt.Println(" Done.")
	}

	p := &pkg{
		name:      name,
		prefix:    prefix,
		suffix:    suffix,
		images:    imgs,
		commands:  cmds,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := p.build(manifest); err != nil {
		p.destroy()
		return errors.Wrap(err, "build package failed")
	}

	fmt.Printf("Successfully built package '%s'\n", name)

	for _, c := range cmds {
		if err := cmd(prefix + c + suffix).create(path.PackageCommand(name, c)); err != nil {
			p.destroy()
			return errors.Wrapf(err, "create command %s", c)
		}
		fmt.Printf("Successfully created command '%s'\n", c)
	}

	fmt.Printf("Successfully installed '%s'.\n", name)

	return nil
}
