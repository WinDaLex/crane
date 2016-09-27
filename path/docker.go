package path

import "path/filepath"

func DockerHost() string {
	return filepath.Join(CranePath(), "var/run/docker.sock")
}

func DockerDaemonGraph() string {
	return filepath.Join(CranePath(), "var/lib/docker")
}

func DockerDaemonPidfile() string {
	return filepath.Join(CranePath(), "var/run/docker.pid")
}
