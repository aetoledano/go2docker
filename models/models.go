package models

import (
	"github.com/aetoledano/go2docker/constants"
	"github.com/pkg/errors"
	"os"
	"regexp"
)

type DkrConfig struct {
	App struct {
		Name string
	}
	Go struct {
		Version string
	}
	IncludeExternalResources []string `yaml:"include-external-resources"`
}

func (config *DkrConfig) Validate() error {
	if len(config.App.Name) == 0 {
		config.App.Name = defaultAppName()
	}

	matched, _ := regexp.MatchString(constants.IMAGE_NAME_REGEX, config.App.Name)
	if !matched {
		return errors.New("Invalid provided docker image name " + config.App.Name)
	}

	if len(config.Go.Version) == 0 {
		config.Go.Version = defaultGoVersion()
	}

	if len(config.IncludeExternalResources) != 0 {
		dir, _ := os.Getwd()
		dir += string(os.PathSeparator)
		for _, name := range config.IncludeExternalResources {
			_, e := os.Stat(dir + name)
			if e != nil {
				return errors.New("Could not access external resource " + e.Error())
			}
		}
	}

	return nil
}

func defaultAppName() string {
	dir, _ := os.Getwd()
	info, _ := os.Stat(dir)
	return info.Name()
}

func defaultGoVersion() string {
	return constants.LATEST_GO_IMAGE_VERSION
}
