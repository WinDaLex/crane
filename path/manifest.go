package path

import "path/filepath"

func Manifests() string {
	return filepath.Join(CranePath(), "manifests")
}

func CraneManifests() string {
	return filepath.Join(Manifests(), "crane")
}

func UserManifests() string {
	return filepath.Join(Manifests(), "user")
}

func UserManifest(pkg string) string {
	return filepath.Join(UserManifests(), pkg+".yml")
}

func RepoManifest(repo string, pkg string) string {
	return filepath.Join(Manifests(), repo, pkg)
}
