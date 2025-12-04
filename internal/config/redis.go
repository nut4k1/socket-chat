package config

func (cfg *Config) RedisAddr() string     { return cfg.Redis.Addr }
func (cfg *Config) RedisPassword() string { return cfg.Redis.Password }
func (cfg *Config) RedisDB() int          { return cfg.Redis.DB }
