package history

import (
	"os"
	"testing"
)

func TestHistory(t *testing.T) {
	os.Setenv("CRANE_PATH", "testdata/crane")

	if Show() != "" {
		t.Fail()
	}

	Add("test")

	if Show() == "" {
		t.Fail()
	}

	Clear()
	if Show() != "" {
		t.Fail()
	}
}
