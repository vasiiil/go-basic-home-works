package config

import (
	"json-bin/output"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key        string
	XMasterKey string
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
	xMasterKey := os.Getenv("X_MASTER_KEY")
	if key == "" {
		panic("Не передан параметр X_MASTER_KEY в переменные окружения")
	}
	return &Config{
		Key:        key,
		XMasterKey: xMasterKey,
	}
}
