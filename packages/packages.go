package packages

import (
	"io/ioutil"
	"os"

	"github.com/windalex/crane/path"
)

func All() ([]string, error) {
	files, _ := ioutil.ReadDir(path.Packages())

	pkgs := []string{}
	for _, f := range files {
		pkgs = append(pkgs, f.Name())
	}
	return pkgs, nil
}

func Exists(pkg string) bool {
	_, err := os.Stat(path.Package(pkg))
	return !os.IsNotExist(err)
}
