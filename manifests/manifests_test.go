package manifests

import (
	"os"
	"testing"
)

func TestInferPackageName(t *testing.T) {
	name, err := InferPackageName("testdata/example")
	if err != nil {
		t.Fatal()
	}

	if name != "example" {
		t.Errorf("want %s, got %s", "example", name)
	}

	name, err = InferPackageName("testdata/example/docker-compose.yml")
	if err != nil {
		t.Fatal()
	}

	if name != "example" {
		t.Errorf("want %s, got %s", "example", name)
	}
}

func TestManifests(t *testing.T) {
	os.Setenv("CRANE_PATH", "testdata/crane")
	if err := Import("testdata/example/docker-compose.yml", "test"); err != nil {
		t.Fatal()
	}

	if Retrive("test") != "testdata/crane/manifests/user/test.yml" {
		t.Errorf(Retrive("test"))
	}

	if err := Export("test"); err != nil {
		t.Fatal()
	}

	if _, err := os.Stat("./test.yml"); os.IsNotExist(err) {
		t.Fail()
	}

	os.Remove("testdata/crane/manifests/user/test.yml")
	os.Remove("./test.yml")
}

func TestUpdate(t *testing.T) {
	os.Unsetenv("CRANE_PATH")
	if err := Update(); err != nil {
		t.Fail()
	}
}
