package path

import "path/filepath"

func History() string {
	return filepath.Join(CranePath(), "history")
}
