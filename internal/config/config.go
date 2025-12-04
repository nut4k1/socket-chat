package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	Auth struct {
		JWTSecret string `yaml:"jwt_secret"`
	} `yaml:"auth"`

	WS struct {
		ReadBufferSize  int `yaml:"read_buffer_size"`
		WriteBufferSize int `yaml:"write_buffer_size"`
		// PingIntervalSec int `yaml:"ping_interval_sec"`
	} `yaml:"ws"`
}

func Load(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	overrideByEnv(&cfg)
	log.Println("Config loaded successfully")

	return &cfg
}

func overrideByEnv(cfg *Config) {
	if env_jwt_secret := os.Getenv("JWT_SECRET"); env_jwt_secret != "" {
		cfg.Auth.JWTSecret = env_jwt_secret
	}

	if env_redis_pw := os.Getenv("REDIS_PW"); env_redis_pw != "" {
		cfg.Redis.Password = env_redis_pw
	}
}
