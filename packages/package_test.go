package packages

import (
	"os"
	"testing"

	"github.com/windalex/crane/path"
)

func TestPackage(t *testing.T) {
	if err := os.Setenv("CRANE_PATH", "testdata/crane"); err != nil {
		t.Fatal()
	}

	if err := Install("hello-world", "testdata/hello-world.yml", "", ""); err != nil {
		t.Fatal()
	}

	if _, err := os.Stat(path.Package("hello-world")); os.IsNotExist(err) {
		t.Fail()
	}

	if !Exists("hello-world") {
		t.Fail()
	}

	if ps, err := All(); err != nil {
		t.Fatal()
	} else {
		if len(ps) != 1 || ps[0] != "hello-world" {
			t.Fail()
		}
	}

	if err := Upgrade("hello-world"); err != nil {
		t.Fail()
	}

	if err := Uninstall("hello-world"); err != nil {
		t.Fatal()
	}

	if _, err := os.Stat(path.Command("hi")); !os.IsNotExist(err) {
		t.Fail()
	}

	if _, err := os.Stat(path.Package("hello-world")); !os.IsNotExist(err) {
		t.Fail()
	}
}

func TestUninstallNonExistingPackage(t *testing.T) {
	if err := Uninstall("pkg-no-exists"); err == nil {
		t.Fail()
	}
}

func TestUpgradeNonExistingPackage(t *testing.T) {
	if err := Upgrade("pkg-no-exists"); err == nil {
		t.Fail()
	}
}

func TestInvalidManifestInstall(t *testing.T) {
	if err := Install("name", "non-existing", "", ""); err == nil {
		t.Fail()
	}
}
