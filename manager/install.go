package manager

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/windalex/crane/history"
	"github.com/windalex/crane/manifests"
	"github.com/windalex/crane/packages"
)

func Install(pkg string, manifest string, prefix string, suffix string) error {
	if pkg == "" && manifest == "" {
		panic("pkg and manifest path should not be empty at the same time")
	}

	if pkg == "" {
		name, err := manifests.InferPackageName(manifest)
		if err != nil {
			return errors.Wrap(err, "infer package name")
		}
		pkg = name
	} else if manifest == "" {
		manifest = manifests.Retrive(pkg)
	}

	if packages.Exists(pkg) {
		return errors.Errorf("%s has already been installed", pkg)
	}

	if err := packages.Install(pkg, manifest, prefix, suffix); err != nil {
		return err
	}

	history.Add(fmt.Sprintf("install %s", pkg))

	return nil
}
