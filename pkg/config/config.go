package config

import (
	"context"

	"github.com/spf13/cobra"
)

type ctxKey int

const cfgKey ctxKey = 1

type Config struct {
	OpenAI     OpenAIConfig `mapstructure:"openai"`
	HttpConfig HttpConfig   `mapstructure:"http"`
}

type OpenAIConfig struct {
	APIKey  string `mapstructure:"api_key"`
	BaseUrl string `mapstructure:"base_url"`
	Model   string `mapstructure:"model"`
}

type HttpConfig struct {
	Proxy *string `mapstructure:"proxy"`
}

func GetCfg(cmd *cobra.Command) *Config {
	return cmd.Context().Value(cfgKey).(*Config)
}

func SetContext(cmd *cobra.Command, cfg *Config) {
	ctx := context.WithValue(cmd.Context(), cfgKey, cfg)
	cmd.SetContext(ctx)
}
