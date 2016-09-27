package packages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/errors"
	"github.com/windalex/crane/path"
)

type pkg struct {
	name      string    `json:"-"`
	prefix    string    `json:"prefix"`
	suffix    string    `json:"suffix"`
	images    []string  `json:"images"`
	commands  []string  `json:"-"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
}

func (p pkg) build(manifest string) error {
	os.Mkdir(path.Package(p.name), 0644)

	exec.Command("cp", manifest, path.PackageComposefile(p.name)).Run()

	// Generate commands

	os.Mkdir(path.PackageCommands(p.name), 0644)
	for _, c := range p.commands {
		script := fmt.Sprintf("#!/bin/sh\ndocker-compose -H unix://%s -f %v run %v %v $*", path.DockerHost(), path.PackageComposefile(p.name), c, c)
		ioutil.WriteFile(path.PackageCommand(p.name, c), []byte(script), 0744)
	}

	// Save metadata in json

	data, _ := json.Marshal(p)
	ioutil.WriteFile(path.PackageMetadata(p.name), data, 0644)

	return nil
}

func retrive(name string) (*pkg, error) {
	p := &pkg{name: name}

	data, err := ioutil.ReadFile(path.PackageMetadata(name))
	if err != nil {
		return nil, errors.Wrapf(err, "read file %s", path.PackageMetadata(name))
	}

	json.Unmarshal(data, p)

	// Load commands

	files, _ := ioutil.ReadDir(path.PackageCommands(name))
	commands := []string{}
	for _, f := range files {
		commands = append(commands, f.Name())
	}

	p.commands = commands

	return p, nil
}

func (p pkg) destroy() {
	os.RemoveAll(path.Package(p.name))
}
