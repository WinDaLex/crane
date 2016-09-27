package packages

import (
	"github.com/pkg/errors"
	"time"
)

func Upgrade(name string) error {
	// Retrive the package, if not found, package is not installed
	p, err := retrive(name)
	if err != nil {
		return errors.Errorf("package is not installed")
	}

	// Pull all images of the package
	for _, i := range p.images {
		if err := img(i).pull(); err != nil {
			return errors.Wrapf(err, "pull image %s", i)
		}
	}

	// Update the update time of the package
	p.updatedAt = time.Now()

	return nil
}
