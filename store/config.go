package store

// Config ...
type Config struct {
	DatabaseUrl string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
