package api

import "json-bin/config"

type Api struct {
	config *config.Config
}

func New(config *config.Config) *Api {
	return &Api{
		config: config,
	}
}