package packages

import (
	"testing"
)

func TestFormatError(t *testing.T) {
	c := composefile("testdata/composefile-format-error.yml")
	if _, _, err := c.validate(); err == nil {
		t.Errorf("composefile validate should return error")
	}
}

func TestNoImages(t *testing.T) {
	c := composefile("testdata/composefile/contains-build.yml")
	if _, _, err := c.validate(); err == nil {
		t.Fail()
	}
}
