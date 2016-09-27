package path

import (
	"path/filepath"
)

func Packages() string {
	return filepath.Join(CranePath(), "packages")
}

func Package(carton string) string {
	return filepath.Join(Packages(), carton)
}

func PackageComposefile(carton string) string {
	return filepath.Join(Package(carton), "docker-compose.yml")
}

func PackageCommands(carton string) string {
	return filepath.Join(Package(carton), "commands")
}

func PackageCommand(carton string, command string) string {
	return filepath.Join(PackageCommands(carton), command)
}

func PackageMetadata(carton string) string {
	return filepath.Join(Package(carton), "metadata.json")
}
