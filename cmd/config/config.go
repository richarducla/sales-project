package config

import (
	"github.com/joeshaw/envdecode"
)

func FromEnv() (Config, error) {
	c := Config{}
	err := envdecode.Decode(&c)
	return c, err
}

type (
	Config struct {
		DbHost     string `env:"DB_HOST,default=localhost"`
		DbPort     string `env:"DB_PORT,default=5432"`
		DbUser     string `env:"DB_USER,default=root"`
		DbPassword string `env:"DB_PASSWORD,default=secret"`
		DbName     string `env:"DB_NAME,default=store"`
		ChunkSize  int    `env:"CHUNK_SIZE,default=4000000"`
		FileName   string `env:"FILE_NAME,default=stock"`
	}
)
