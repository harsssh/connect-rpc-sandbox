package config

type Config struct {
	Host string
	Port string
}

func GetConfig() *Config {
	// TODO: env を読み込む
	return &Config{
		Host: "127.0.0.1",
		Port: "8080",
	}
}
