package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		MaxFirstHour        int `mapstructure:"maxFirstHour" validate:"required"`
		ChargeFirstTwoHours int `mapstructure:"chargeFirstTwoHours" validate:"required"`
		ChargeNextHours     int `mapstructure:"chargeNextHours" validate:"required"`
	}
)

func (c *Config) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	return err
}

func (c *Config) LoadDefaultValue(loader *viper.Viper) {
	loader.SetDefault("config.chargeFirstTwoHours", "10")
	loader.SetDefault("config.chargeNextHours", "10")
}
