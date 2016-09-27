package manager

import (
	"github.com/pkg/errors"
	"github.com/windalex/crane/packages"
)

func UpgradePackage(pkg string) error {
	return packages.Upgrade(pkg)
}

func UpgradeAllPackages() error {
	pkgs, err := packages.All()
	if err != nil {
		return errors.Wrap(err, "get all installed packages error")
	}

	for _, pkg := range pkgs {
		if err := packages.Upgrade(pkg); err != nil {
			return err
		}
	}

	return nil
}
