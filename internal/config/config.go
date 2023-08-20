package config

type Config struct {
	path *string
}

func NewConfig(path *string) *Config {
	return &Config{
		path: path,
	}
}

func (c *Config) Read() []string {
	return []string{
		"url1",
		"url2",
		"url3",
	}
}
