package manifests

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/windalex/crane/path"
)

func InferPackageName(manifest string) (string, error) {
	if filepath.Base(manifest) == "docker-compose.yml" || filepath.Base(manifest) == "docker-compose.yaml" {
		abspath, _ := filepath.Abs(manifest)
		return filepath.Base(filepath.Dir(abspath)), nil
	}
	return strings.TrimSuffix(filepath.Base(manifest), filepath.Ext(manifest)), nil
}

func All() (map[string]string, error) {
	dirs := []string{"user", "crane"}

	manifests := map[string]string{}
	for _, dir := range dirs {
		p := filepath.Join(path.Manifests(), dir)
		files, err := ioutil.ReadDir(p)
		if err != nil {
			return nil, errors.Wrap(err, "read packages dir failed")
		}

		for _, f := range files {
			filename := f.Name()
			manifest := strings.TrimSuffix(filename, filepath.Ext(filename))
			if _, prs := manifests[manifest]; !prs {
				manifests[manifest] = path.RepoManifest(dir, filename)
			}
		}
	}

	return manifests, nil
}

func Retrive(pkg string) string {
	l, _ := All()
	return l[pkg]
}

func Update() error {
	os.Chdir(path.CraneManifests())
	if err := exec.Command("git", "pull").Run(); err != nil {
		return errors.Errorf("git pull failed")
	}
	return nil
}

func Import(manifest string, pkg string) error {
	pwd, _ := os.Getwd()
	if err := exec.Command("cp", filepath.Join(pwd, manifest), path.UserManifest(pkg)).Run(); err != nil {
		return err
	}

	return nil
}

func Export(pkg string) error {
	if err := exec.Command("mv", path.UserManifest(pkg), filepath.Join(os.Getenv("PWD"), pkg+".yml")).Run(); err != nil {
		return err
	}
	return nil
}
