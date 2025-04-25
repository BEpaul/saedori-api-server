package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Server struct {
		Port string `toml:"port"`
		CrawlApiBaseUrl string `toml:"crawl_api_base_url"`
	}
}

func NewConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	
	filePath := fmt.Sprintf("config.%s.toml", env)
	
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	return c
}
