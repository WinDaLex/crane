package packages

import (
	"os"

	"github.com/pkg/errors"
	"github.com/windalex/crane/path"
)

type cmd string

func (c cmd) exists() bool {
	_, err := os.Readlink(c.path())
	return err == nil
}

func (c cmd) create(src string) error {
	if err := os.Symlink(src, c.path()); err != nil {
		return errors.Wrapf(err, "create symbolic link (%s to %s)", c.path(), src)
	}
	return nil
}

func (c cmd) destroy() error {
	if err := os.Remove(c.path()); err != nil {
		return errors.Wrapf(err, "remove %s", c.path())
	}
	return nil
}

func (c cmd) path() string {
	return path.Command(string(c))
}
