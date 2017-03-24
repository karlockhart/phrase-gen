package config

import "github.com/spf13/viper"
import "log"

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/phrase-gen/")
	viper.ReadInConfig()
	log.Println("Config")
}
