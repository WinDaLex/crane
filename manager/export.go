package manager

import (
	"github.com/windalex/crane/manifests"
)

func Export(pkg string) error {
	return manifests.Export(pkg)
}
