package manager

import (
	"github.com/windalex/crane/history"
)

func ShowHistory() string {
	return history.Show()
}

func ClearHistory() {
	history.Clear()
}
