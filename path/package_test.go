package path

import (
	"os"
	"testing"
)

func init() {
	if err := os.Setenv("CRANE_PATH", "test"); err != nil {
		panic(`os.Setenv("CRANE_PATH", "test") error`)
	}
}

func TestPackages(t *testing.T) {
	if Packages() != "test/packages" {
		t.Fail()
	}
}

func TestPackage(t *testing.T) {
	if Package("carton") != "test/packages/carton" {
		t.Fail()
	}
}

func TestPackageComposefile(t *testing.T) {
	if PackageComposefile("carton") != "test/packages/carton/docker-compose.yml" {
		t.Fail()
	}
}

func TestPackageCommands(t *testing.T) {
	if PackageCommands("carton") != "test/packages/carton/commands" {
		t.Fail()
	}
}

func TestPackageCommand(t *testing.T) {
	if PackageCommand("carton", "command") != "test/packages/carton/commands/command" {
		t.Fail()
	}
}

func TestMetadata(t *testing.T) {
	if PackageMetadata("carton") != "test/packages/carton/metadata.json" {
		t.Fail()
	}
}
