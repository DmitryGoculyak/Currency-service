package config

import (
	"Currency-service/internal/db"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type GrpcServiceConfig struct {
	Host string
	Port string
}

var (
	err    error
	config *Config
	s      sync.Once
)

type Config struct {
	DBConfig   *db.DBConfig
	GrpcConfig *GrpcServiceConfig
}

func LoadConfig() (*Config, error) {
	s.Do(func() {
		config = &Config{}

		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		DBConfig := viper.Sub("database")
		GrpcConfig := viper.Sub("service")

		if err = parseSubConfig(DBConfig, &config.DBConfig); err != nil {
			return
		}
		if err = parseSubConfig(GrpcConfig, &config.GrpcConfig); err != nil {
			return
		}
	})
	return config, err
}

func parseSubConfig[T any](subConfig *viper.Viper, parseTo *T) error {
	if subConfig == nil {
		return fmt.Errorf("can not read %T config: subconfig is nil", parseTo)
	}

	if err = subConfig.Unmarshal(parseTo); err != nil {
		return err
	}
	return nil
}
