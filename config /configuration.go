package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type AppConfiguration struct {
	ENV              string
	DBDebug          bool
	ApiPrefix        string
	ApiKey           string
	TestingApiKey    string
	SuffixForTracing string
}

type ServerConfiguration struct {
	Port int
}

type DatabaseConfiguration struct {
	Driver         string
	Name           string
	User           string
	Password       string
	Host           string
	Port           string
	MigrationsPath string
}

// Configuration config
type Configuration struct {
	App      AppConfiguration
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

var (
	configuration *Configuration
	once          sync.Once
)

// All get all config
func All() *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("${PWD}/.")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AllowEmptyEnv(true)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})

	return configuration
}

func Get() *Configuration {
	return configuration
}

func GetSuffixForTracing() string {
	if configuration == nil {
		return ""
	}

	return configuration.App.SuffixForTracing
}

// for testing purpose
func Set(conf *Configuration) {
	configuration = conf
}
