package bootstrap

import (
	"github.com/roniahmad/parking-app/config"
	"github.com/roniahmad/parking-app/internal/configloader"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewConfig() *config.Config {
	configPath := "./config.yaml"

	loader := configloader.New(configPath,
		// before config loaded hook
		func(c *config.Config, loader *viper.Viper) {
			c.LoadDefaultValue(loader)
		},
		// after config loaded hooks
		func(c *config.Config, loader *viper.Viper) {
			// validate configuration
			if err := c.Validate(); err != nil {
				log.Fatal().Msgf("Error validating configuration: %s", err)
			}
		},
	)

	// get config and expand conf variable to configuration array
	conf := loader.GetConfig()

	return &conf
}
