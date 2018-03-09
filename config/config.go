package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init () {
	env := os.Getenv("GO_ENV")
	config := viper.New()
	config.AddConfigPath("./config")

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if env == "" || env == "development" {
		c := config.GetStringMapString("development")
		for key, value := range c {
			os.Setenv(strings.ToUpper(key), value)
		}
	}

}