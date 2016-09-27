package path

import "path/filepath"

func Commands() string {
	return filepath.Join(CranePath(), "bin")
}

func Command(command string) string {
	return filepath.Join(Commands(), command)
}
