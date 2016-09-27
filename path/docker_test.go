package path

import (
	"os"
	"testing"
)

func init() {
	if err := os.Setenv("CRANE_PATH", "test"); err != nil {
		panic(`os.Setenv("CRANE_PATH", "test") error`)
	}
}

func TestDockerHost(t *testing.T) {
	if DockerHost() != "test/var/run/docker.sock" {
		t.Fail()
	}
}

func TestDockerDaemonGraph(t *testing.T) {
	if DockerDaemonGraph() != "test/var/lib/docker" {
		t.Fail()
	}
}

func TestDockerDaemonPidfile(t *testing.T) {
	if DockerDaemonPidfile() != "test/var/run/docker.pid" {
		t.Fail()
	}
}
