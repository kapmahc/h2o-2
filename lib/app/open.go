package app

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Read read from config file
func Read(file string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Write write config file
func Write(file string, cfg *Config) error {
	fd, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()

	enc := toml.NewEncoder(fd)
	return enc.Encode(cfg)
}
