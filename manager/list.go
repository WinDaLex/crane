package manager

import (
	"fmt"

	"github.com/windalex/crane/manifests"
	"github.com/windalex/crane/packages"
)

func AllInstalledPackages() ([]string, error) {
	return packages.All()
}

func AllManifests() ([]string, error) {
	r, _ := manifests.All()
	fmt.Println(r)
	return []string{}, nil
}
