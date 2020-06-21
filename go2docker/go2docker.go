package go2docker

import (
	"io/ioutil"
	"github.com/aetoledano/go2docker/constants"
)

func Dockerizeit() error {
	rawGo2dockerConfig, err := ioutil.ReadFile(GO2DOCKER_FILE)
	if err != nil {
		return err
	}

	return nil
}
