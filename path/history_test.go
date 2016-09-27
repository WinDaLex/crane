package path

import (
	"os"
	"testing"
)

var test = "test"

func TestHistory(t *testing.T) {
	if err := os.Setenv("CRANE_PATH", test); err != nil {
		t.Fatal("setenv error")
	}

	want := test + "/history"
	got := History()
	if got != want {
		t.Error("want", want, "got", got)
	}
}
