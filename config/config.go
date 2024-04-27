package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var file = "config.yaml" // 配置文件名

type Config struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`
	Redis struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"redis"`
	System struct {
		Name string `yaml:"server_name"`
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"system"`
}

func InitConfig() (Config, error) {
	var config Config
	args := os.Args
	if len(args) > 1 {
		file = args[1]
	}

	// 读取并解析YAML文件
	yamlData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Failed to read YAML file: %v", err)
		return config, err
	}

	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML: %v", err)
		return config, err
	}

	// 打印解析后的数据
	return config, nil
}
