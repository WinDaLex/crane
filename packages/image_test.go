package packages

import "testing"

func TestImageName(t *testing.T) {
	i := img("example")

	if string(i) != "example" {
		t.Errorf("string(i) expected \"example\", got \"%s\"", string(i))
	}
}

func TestImagePullAndRemove(t *testing.T) {
	i := img("hello-world")

	if err := i.pull(); err != nil {
		t.Errorf("image pull error")
	}

	if err := i.remove(); err != nil {
		t.Errorf("image remove error")
	}
}

func TestImagePullError(t *testing.T) {
	i := img("image-not-exist")

	if err := i.pull(); err == nil {
		t.Errorf("image pull should return error")
	}
}

func TestImageRemoveError(t *testing.T) {
	i := img("image-not-pulled-yet")

	if err := i.remove(); err == nil {
		t.Errorf("image remove should return error")
	}
}
