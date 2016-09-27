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

func TestCommands(t *testing.T) {
	if Commands() != "test/bin" {
		t.Fail()
	}
}

func TestCommand(t *testing.T) {
	if Command("command") != "test/bin/command" {
		t.Fail()
	}
}
