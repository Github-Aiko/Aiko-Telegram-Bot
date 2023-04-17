package config

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var (
	configFile = flag.String("config", "", "Config file for Aiko Bot Telegram.")
)

func GetConfig() *viper.Viper {
	config := viper.New()

	// Set custom path and name
	if *configFile != "" {
		configName := path.Base(*configFile)
		configFileExt := path.Ext(*configFile)
		configNameOnly := strings.TrimSuffix(configName, configFileExt)
		configPath := path.Dir(*configFile)
		config.SetConfigName(configNameOnly)
		config.SetConfigType(strings.TrimPrefix(configFileExt, "."))
		config.AddConfigPath(configPath)
		// Set ASSET Path and Config Path
		os.Setenv("ASSET_PATH", configPath)
		os.Setenv("CONFIG_PATH", configPath)
	} else {
		// Set default config path
		config.SetConfigName("config")
		config.SetConfigType("yaml")
		config.AddConfigPath(".")
	}

	if err := config.ReadInConfig(); err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	config.WatchConfig()
	return config
}
