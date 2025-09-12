package config

import (
	"crypto/rand"
	"encoding/hex"
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

	Auth struct {
		JWTSecret string `yaml:"jwt_secret"`
		TokenTTL  int    `yaml:"token_ttl"`
	} `yaml:"auth"`

	AdminUser string `yaml:"admin_user"`

	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
}

var Cfg *Config

func defaultConfig() *Config {
	cfg := &Config{}

	cfg.Server.Host = "0.0.0.0"
	cfg.Server.Port = 8080

	cfg.Auth.JWTSecret = randomJWTSecret(32)
	cfg.Auth.TokenTTL = 3600

	cfg.AdminUser = "admin"

	//cfg.Logging.Level = "warn"
	cfg.Logging.Level = "debug"
	cfg.Logging.File = "app.log"

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

func randomJWTSecret(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("не удалось сгенерировать JWT секрет: %v", err)
	}
	return hex.EncodeToString(b)
}
