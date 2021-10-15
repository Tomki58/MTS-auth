package config

// A Config is a struct for configuration
type Config struct {
	Debug bool
}

// Cfg contains dynamic info about service configuration
var Cfg *Config

func init() {
	Cfg = new(Config)
}

// New returns the new Config instance
func New() *Config {
	return new(Config)
}

// Switch changes the config's debug value
func (c *Config) Switch() {
	c.Debug = !c.Debug
}
