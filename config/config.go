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
	if env == "" {
		env = "development"
	}
	if env == "development" || env == "test" {
		c := config.GetStringMapString(env)
		for key, value := range c {
			os.Setenv(strings.ToUpper(key), value)
		}
	}

}