package config

import (
	"json-bin/output"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не загрузились env")
	}
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	if !(len(key) == 16 || len(key) == 24 || len(key) == 32) {
		panic("Некорректная длина ключа")
	}
	return &Config{
		Key: key,
	}
}
