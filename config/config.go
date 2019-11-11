package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func (databaseConfig DatabaseConfig) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Database)
}

func SetupDatabaseConfig(databaseConfigPath string) (DatabaseConfig, error) {
	var databaseConfig DatabaseConfig

	databaseConfigFile, err := ioutil.ReadFile(databaseConfigPath)
	if err != nil {
		return databaseConfig, err
	}

	err = yaml.Unmarshal(databaseConfigFile, &databaseConfig)
	if err != nil {
		return databaseConfig, err
	}

	return databaseConfig, err
}
