package broker

type FakeRedisConfig struct {
	Addr string
	Pass string
	DB   int
}

func (frc *FakeRedisConfig) RedisAddr() string     { return frc.Addr }
func (frc *FakeRedisConfig) RedisPassword() string { return frc.Pass }
func (frc *FakeRedisConfig) RedisDB() int          { return frc.DB }
