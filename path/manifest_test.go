package path

import "testing"

func TestCraneManifests(t *testing.T) {
	if CraneManifests() != CranePath()+"/manifests/crane" {
		t.Fail()
	}
}

func TestUserManifest(t *testing.T) {
	expected := CranePath() + "/manifests/user/example.yml"
	got := UserManifest("example")
	if got != expected {
		t.Errorf("want %s, got%s\n", expected, got)
	}
}

func TestRepoManifest(t *testing.T) {
	if RepoManifest("repo", "pkg") != Manifests()+"/repo/pkg" {
		t.Fail()
	}
}
