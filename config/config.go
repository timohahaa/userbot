package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server   `yaml:"server"`
		DB       `yaml:"database"`
		Telegram `yaml:"telegram"`
	}
	Server struct {
		Port string `yaml:"port" env:"HTTP_SERVER_PORT"`
	}
	DB struct {
		URL          string `yaml:"url" env:"PG_URL" env-required:"true"`
		ConnPoolSize int    `yaml:"maxConnPoolSize" env:"PG_MAX_POOL_SIZE"`
	}
	Telegram struct {
		ApiId   int    `yaml:"API_ID"`
		ApiHash string `yaml:"API_HASH"`
		Bot     `yaml:"bot"`
	}
	Bot struct {
		CommentFrequency int `yaml:"commentFrequency"`
	}
)

func NewConfig(filePath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(filePath, cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	//	err = cleanenv.UpdateEnv(cfg)
	//	if err != nil {
	//		return nil, fmt.Errorf("error updating env: %w", err)
	//	}

	return cfg, nil
}
