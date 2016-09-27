package path

import (
	"os"
	"testing"
)

func TestCranePath(t *testing.T) {
	testCranePath := "test_path"

	if err := os.Unsetenv("CRANE_PATH"); err != nil {
		t.Fatal(`Unsetenv("CRANE_PATH") error`)
	}

	if CranePath() != os.Getenv("HOME")+"/.crane" {
		t.Error("want", os.Getenv("HOME")+"/.crane", "got", CranePath())
	}

	if err := os.Setenv("CRANE_PATH", testCranePath); err != nil {
		t.Fatal("setenv error")
	}

	if CranePath() != testCranePath {
		t.Errorf("want %s, got %s", testCranePath, CranePath())
	}
}
