package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (databaseConfig DatabaseConfig) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Database)
}

func SetupConfig() (DatabaseConfig, error) {
	var databaseConfig DatabaseConfig

	if os.Getenv("USERNAME_DATABASE") != "" {
		databaseConfig.Username = os.Getenv("USERNAME_DATABASE")
	}

	if os.Getenv("PASSWORD_DATABASE") != "" {
		databaseConfig.Password = os.Getenv("PASSWORD_DATABASE")
	}

	if os.Getenv("HOST_DATABASE") != "" {
		databaseConfig.Host = os.Getenv("HOST_DATABASE")
	} else {
		databaseConfig.Host = "localhost"
	}

	if os.Getenv("PORT_DATABASE") != "" {
		databaseConfig.Port = os.Getenv("PORT_DATABASE")
	}

	if os.Getenv("DATABASE_NAME") != "" {
		databaseConfig.Database = os.Getenv("DATABASE_NAME")
	}

	return databaseConfig, nil
}
