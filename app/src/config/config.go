package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		APIServer `mapstructure:",squash" validate:"required"`
		RDB       `mapstructure:",squash" validate:"required"`
	}

	APIServer struct {
		APIPort string `mapstructure:"API_PORT" validate:"required"`
		APIHost string `mapstructure:"API_HOST" validate:"required"`
	}

	RDB struct {
		DBMS   string `mapstructure:"DB_MS" validate:"required"`
		DBName string `mapstructure:"DB_NAME" validate:"required"`
		DBHost string `mapstructure:"DB_HOST" validate:"required"`
		DBPort string `mapstructure:"DB_PORT" validate:"required"`
		DBUser string `mapstructure:"DB_USER" validate:"required"`
		DBPass string `mapstructure:"DB_PASS" validate:"required"`
	}
)

func NewConfig(confPath string) (*Config, error) {
	c := new(Config)

	viper.SetConfigType("env")
	viper.SetConfigFile(confPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	validator := validator.New()
	if err := validator.Struct(c); err != nil {
		return nil, err
	}

	return c, nil
}
