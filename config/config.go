package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	AppVersion string `mapstructure:"VERSION"`
	Host       string `mapstructure:"HOST"`
	Port       int    `mapstructure:"PORT"`
	Debug      bool   `mapstructure:"DEBUG"`
}

type PostgresConfig struct {
	PostgresURL      string `mapstructure:"URL"`
	PostgresHost     string `mapstructure:"HOST"`
	PostgresPort     int    `mapstructure:"PORT"`
	PostgresUser     string `mapstructure:"USER"`
	PostgresPassword string `mapstructure:"PASSWORD"`
	PostgresDB       string `mapstructure:"DB"`
	PostgresPgDriver string `mapstructure:"PG_DRIVER"`
}

func LoadConfig(path string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
