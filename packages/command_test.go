package packages

import (
	"os"
	"testing"
)

func TestCommand(t *testing.T) {
	c := cmd("example")

	if err := c.create("testdata/launcher"); err != nil {
		t.Fatalf("call CreateCommand error: %v\n", err)
	}

	if !c.exists() {
		t.Errorf("expected command exists but not, or command exist failed")
	}

	if err := c.destroy(); err != nil {
		t.Fatalf("call delete command failed: %v\n", err)
	}

	if c.exists() {
		t.Errorf("expected command removed but not, or command exist failed")
	}
}

func TestCreateError(t *testing.T) {
	c := cmd("example")

	if err := c.create("testdata/launcher"); err != nil {
		t.Fatalf("call CreateCommand error: %v\n", err)
	}

	if err := c.create("testdata/launcher"); err == nil {
		t.Errorf("create should return error")
	}

	if err := os.Remove(c.path()); err != nil {
		t.Fatalf("call CreateCommand error: %v\n", err)
	}
}

func TestDestroyError(t *testing.T) {
	c := cmd("example")

	if err := c.destroy(); err == nil {
		t.Errorf("destroy should return error")
	}
}
