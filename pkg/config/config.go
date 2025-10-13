package config

import "context"

type CtxKey int

const CfgKey CtxKey = 1

type Config struct {
	OpenAI     OpenAIConfig `mapstructure:"openai"`
	HttpConfig HttpConfig   `mapstructure:"http"`
	Console    Console      `mapstructure:"console"`
}

type OpenAIConfig struct {
	APIKey  string `mapstructure:"api_key"`
	BaseUrl string `mapstructure:"base_url"`
	Model   string `mapstructure:"model"`
}

type HttpConfig struct {
	Proxy *string `mapstructure:"proxy"`
}

type Console struct {
	Output string `mapstructure:"output"`
}

func GetCfg(ctx context.Context) *Config {
	return ctx.Value(CfgKey).(*Config)
}

func SetCfg(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, CfgKey, cfg)
}
