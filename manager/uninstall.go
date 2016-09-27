package manager

import (
	"fmt"

	"github.com/windalex/crane/history"
	"github.com/windalex/crane/packages"
)

func Uninstall(pkg string) error {
	fmt.Printf("uninstall %s...\n", pkg)

	if err := packages.Uninstall(pkg); err != nil {
		return err
	}

	history.Add(fmt.Sprintf("uninstall %s", pkg))

	return nil
}
