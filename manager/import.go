package manager

import (
	"github.com/windalex/crane/manifests"
)

func Import(manifest string) error {
	pkg, err := manifests.InferPackageName(manifest)
	if err != nil {
		return err
	}
	return manifests.Import(manifest, pkg)
}
