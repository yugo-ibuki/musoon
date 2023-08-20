package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct{}

type Content struct {
	ID string `toml:"id"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Read(path string) (Content, error) {
	fmt.Print("Reading the config file...\n")

	var content Content
	if _, err := toml.DecodeFile(path, &content); err != nil {
		return Content{}, err
	}

	fmt.Print("The config file has been read.\n")
	return content, nil
}
