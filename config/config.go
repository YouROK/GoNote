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

	Site struct {
		Host string `yaml:"host"`
	} `yaml:"site"`

	Antispam struct {
		MaxRequests int `yaml:"max_requests"`
		WindowSec   int `yaml:"window_sec"`
	} `yaml:"antispam"`

	Counter struct {
		TTLSeconds int `yaml:"ttl_seconds"` // Время жизни cookie для счётчика
	} `yaml:"counter"`

	TGBot struct {
		Token        string  `yaml:"token"`
		StartMessage string  `yaml:"start_message"`
		AdminIds     []int64 `yaml:"admin_ids"`
		MsgOnNewNote bool    `yaml:"msg_on_new_note"`
	}
}

var Cfg *Config

func defaultConfig() *Config {
	hostname := ""
	if h, err := os.Hostname(); err == nil {
		hostname = h
	} else {
		hostname = "localhost"
	}

	cfg := &Config{}

	cfg.Server.Host = "0.0.0.0"
	cfg.Server.Port = 8095

	cfg.Site.Host = hostname

	cfg.Antispam.MaxRequests = 10
	cfg.Antispam.WindowSec = 30

	cfg.Counter.TTLSeconds = 3600

	cfg.TGBot.Token = ""
	cfg.TGBot.StartMessage = "Hi! 👋\nI'm the GoNote bot for receiving complaint notifications from the website"
	cfg.TGBot.AdminIds = make([]int64, 0)
	cfg.TGBot.MsgOnNewNote = false

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
