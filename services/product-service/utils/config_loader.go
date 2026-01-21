package units

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Type               string `mapstructure:"type"`
		PostgresConnString string `mapstructure:"postgresConnString"` // PostgreSQL
		MysqlConnString    string `mapstructure:"mysqlConnString"`    // MySQL
	} `mapstructure:"database"`
}

func LoadConfig() Config {
	var config Config
	viper.SetConfigFile("/app/pkg/config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshalling config file", err)
	}
	return config
}
