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
}

var Cfg *Config

func defaultConfig() *Config {
	cfg := &Config{}

	cfg.Server.Host = "0.0.0.0"
	cfg.Server.Port = 8095

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
		// создаём файл с настройками по умолчанию
		cfg := defaultConfig()
		if err := SaveConfig(filename, cfg); err != nil {
			log.Fatalf("не удалось создать config.yaml: %v", err)
		}
		log.Printf("создан config.yaml с настройками по умолчанию")
		Cfg = cfg
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("не удалось прочитать файл настроек: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("ошибка разбора настроек: %v", err)
	}

	Cfg = &cfg
}
