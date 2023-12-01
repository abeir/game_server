package conf

import (
	"game_server/server/common/utils"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type YamlConfigLoader struct {
	finder IConfigFinder
}

func (y *YamlConfigLoader) SetFinder(finder IConfigFinder) {
	y.finder = finder
}

func (y *YamlConfigLoader) Load() (ServerConfig, error) {
	conf := ServerConfig{}
	if y.finder == nil {
		return conf, utils.ArgumentError("not set finder")
	}
	filePath, ok := y.finder.Find()
	if !ok {
		return conf, utils.ArgumentError("config not found")
	}
	isExist, _ := utils.PathExists(filePath)
	if !isExist {
		return conf, utils.PathError("config not exists")
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return conf, err
	}
	defer utils.CloseQuiet(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return conf, err
	}

	if err = yaml.Unmarshal(data, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}
