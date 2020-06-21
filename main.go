package main

import (
	"github.com/aetoledano/go2docker/constants"
	app "github.com/aetoledano/go2docker/go2docker"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		panic(constants.NO_TARGET_SUPPLIED)
	}

	cwd := argsWithoutProg[0]

	err := os.Chdir(cwd)
	if err != nil {
		panic(constants.NOT_VALID_TARGET)
	}

	err = app.Dockerizeit()
	if err != nil {
		panic(err.Error())
	}
}
