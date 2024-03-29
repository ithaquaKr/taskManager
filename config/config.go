package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// All config struct
type Config struct {
	App      AppConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

// App config struct
type AppConfig struct {
	AppVersion string `mapstructure:"VERSION"`
	Host       string `mapstructure:"HOST"`
	Port       int    `mapstructure:"PORT"`
	Debug      bool   `mapstructure:"DEBUG"`
	Mode       string `mapstructure:"MODE"`
}

// Postgres config struct
type PostgresConfig struct {
	PostgresURL      string `mapstructure:"URL"`
	PostgresHost     string `mapstructure:"HOST"`
	PostgresPort     int    `mapstructure:"PORT"`
	PostgresUser     string `mapstructure:"USER"`
	PostgresPassword string `mapstructure:"PASSWORD"`
	PostgresDB       string `mapstructure:"DB"`
	PostgresPgDriver string `mapstructure:"PG_DRIVER"`
}

// Logger config struct
type LoggerConfig struct {
	Development       bool   `mapstructure:"DEVELOPMENT"`
	DisableCaller     bool   `mapstructure:"DISABLE_CALLER"`
	DisableStacktrace bool   `mapstructure:"DISABLE_STACKTRACE"`
	Encoding          string `mapstructure:"ENCODING"`
	Level             string `mapstructure:"LEVEL"`
}

// LoadConfig loads configuration from given path
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

// ParseConfig parses config from viper instance
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
