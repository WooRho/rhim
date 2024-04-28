package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var file = "config.yaml" // 配置文件名

type (
	Mysql struct {
		Host              string `yaml:"host"`
		Port              string `yaml:"port"`
		User              string `yaml:"user"`
		Password          string `yaml:"password"`
		Charset           string `yaml:"charset"`
		Collation         string `yaml:"collation"`
		Database          string `yaml:"database"`
		MaxIdleConnection int    `yaml:"max_idle_connection"`
		MaxConnection     int    `yaml:"max_connection"`
	}
	Redis struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
	System struct {
		Name string `yaml:"server_name"`
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
	Config struct {
		Mysql  Mysql  `yaml:"mysql"`
		Redis  Redis  `yaml:"redis"`
		System System `yaml:"system"`
	}
)

var (
	mysql  Mysql
	redis  Redis
	system System
)

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
	mysql = config.Mysql
	redis = config.Redis
	system = config.System
	// 打印解析后的数据
	return config, nil
}

func GetMysql() Mysql {
	return mysql
}

func GetRedis() Redis {
	return redis
}
