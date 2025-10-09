package config

import "time"

type Config struct {
	OpenaiApiKey string        `json:"openai_api_key"`
	TimeOut      time.Duration `json:"timeout"`
	BaseUrl      string        `json:"base_url"`
}
