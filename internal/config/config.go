package config

import "flag"

const serverAddrDefault = ":8080"

type Config struct {
	ServerAddr string
}

func LoadConfig() Config {
	cfg := Config{ServerAddr: serverAddrDefault}

	cfg.loadFlag()

	return cfg
}

func (cfg *Config) loadFlag() {
	flag.StringVar(&cfg.ServerAddr, "addr", cfg.ServerAddr, "HTTP server address")

	flag.Parse()
}
