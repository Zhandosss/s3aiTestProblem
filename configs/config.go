package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"http_server"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func Load() *Config {
	viper.SetDefault("port", "8080")
	viper.SetDefault("host", "localhost")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./build/configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Msgf("failed to read config file: %s", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal().Msgf("failed to unmarshal config: %s", err)
	}

	return config
}
