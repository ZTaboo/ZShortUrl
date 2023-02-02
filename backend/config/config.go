package config

import (
	"ZShortUrl/model"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	Conf *model.Config
)

func init() {
	initConfig()
}

func initConfig() {
	filePath := "config/config.yaml"
	config := new(model.Config)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("open config error:", err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("unmarshal error:", err)
	}
	Conf = config

}
