package go2docker

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aetoledano/go2docker/constants"
	"github.com/aetoledano/go2docker/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/jhoonb/archivex"
	"github.com/valyala/fastjson"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Dockerizeit() error {
	var err error
	var config models.DkrConfig

	rawGo2dockerConfig, err := ioutil.ReadFile(constants.GO2DOCKER_FILE)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rawGo2dockerConfig, &config)
	if err != nil {
		return err
	}

	err = config.Validate()
	if err != nil {
		return err
	}

	dockerfile := createDockerFile(&config)

	target, err := createTarBuildContext(dockerfile)
	if err != nil {
		return err
	}

	buildContext, err := os.Open(*target)
	if err != nil {
		return err
	}
	defer buildContext.Close()

	err = buildDockerImage(buildContext, &config)
	if err != nil {
		return err
	}

	cleanTheHouse(*target)

	return nil
}

func buildDockerImage(buildCtx *os.File, config *models.DkrConfig) error {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	tags := []string{}
	tags = append(tags, config.App.Name)

	opt := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		Dockerfile:     constants.DOCKERFILE,
		Tags:           tags,
	}

	response, err := cli.ImageBuild(context.Background(), buildCtx, opt)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(response.Body)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Print(fastjson.GetString(bytes, "stream"))
	}

	return nil
}

func createDockerFile(config *models.DkrConfig) string {
	str := constants.DOCKERFILE_TEMPLATE

	str = strings.ReplaceAll(str, constants.IMAGE_VERSION, config.Go.Version)
	str = strings.ReplaceAll(str, constants.APP_NAME, config.App.Name)
	str = strings.ReplaceAll(str, constants.EXEC_NAME, uuid.New().String())

	return str
}

func createTarBuildContext(dockerfile string) (*string, error) {

	target := os.TempDir() + string(os.PathSeparator) + uuid.New().String() + constants.CTX_SUFFIX
	source, _ := os.Getwd()

	tarFile := new(archivex.TarFile)
	err := tarFile.Create(target)
	if err != nil {
		return nil, err
	}

	err = tarFile.AddAll(source, false)
	if err != nil {
		return nil, err
	}

	err = tarFile.Add(constants.DOCKERFILE, strings.NewReader(dockerfile), nil)
	if err != nil {
		return nil, err
	}

	err = tarFile.Close()
	if err != nil {
		return nil, err
	}

	return &target, nil
}

func cleanTheHouse(target string) {
	err := os.Remove(target)
	if err != nil {
		fmt.Println("Warn: Couldn't remove tmp build context. " + err.Error())
	}
}
