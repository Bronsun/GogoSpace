package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

const (
	Port               = "PORT"
	ApiKey             = "API_KEY"
	ConcurrentRequests = "CONCURRENT_REQUESTS"

	DefaultPort               = 8080
	DefaultApiKey             = "DEMO_KEY"
	DefaultConcurrentRequests = 5
)

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetPort() string {
	port := fmt.Sprint(config.GetInt(Port))
	return (":" + port)
}

func GetApiKey() string {
	return config.GetString(ApiKey)
}

func GetConcurrentRequests() int {
	return config.GetInt(ConcurrentRequests)
}
