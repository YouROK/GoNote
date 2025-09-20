package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Antispam struct {
		MaxRequests int `yaml:"max_requests"`
		WindowSec   int `yaml:"window_sec"`
	} `yaml:"antispam"`
}

var Cfg *Config

func defaultConfig() *Config {
	cfg := &Config{}

	cfg.Server.Host = "0.0.0.0"
	cfg.Server.Port = 8095

	cfg.Antispam.MaxRequests = 10
	cfg.Antispam.WindowSec = 30

	return cfg
}

func SaveConfig(filename string, cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadConfig() {
	filename := filepath.Join(filepath.Dir(os.Args[0]), "config.yaml")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		cfg := defaultConfig()
		if err := SaveConfig(filename, cfg); err != nil {
			log.Fatalf("failed to create config.yaml: %v", err)
		}
		log.Printf("config.yaml created with default settings")
		Cfg = cfg
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	Cfg = &cfg
	log.Printf("config.yaml loaded successfully")
}
