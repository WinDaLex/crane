package path

import (
	"os"
	"path/filepath"
)

func CranePath() string {
	if os.Getenv("CRANE_PATH") == "" {
		return filepath.Join(os.Getenv("HOME"), ".crane")
	} else {
		return os.Getenv("CRANE_PATH")
	}
}
