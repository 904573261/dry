package docker

import (
	"os"
)

//Env holds Docker-related environment variables
type Env struct {
	DockerHost       string
	DockerTLSVerify  bool //tls must be verified
	DockerCertPath   string
	DockerAPIVersion string
}

//NewEnv creates a new docker environment struct
func NewEnv() *Env {
	version := os.Getenv("DOCKER_API_VERSION")
	if version == "" {
		//version = api.DefaultVersion
		version = "1.37"

	}
	return &Env{DockerAPIVersion: version}
}
