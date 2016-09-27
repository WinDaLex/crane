package manager

import (
	"github.com/windalex/crane/manifests"
)

func UpdateManifests() error {
	return manifests.Update()
}
