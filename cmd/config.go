package cmd

type Config struct {
	OpenAI     OpenAIConfig `mapstructure:"openai"`
	HttpConfig HttpConfig   `mapstructure:"http"`
}

type OpenAIConfig struct {
	APIKey string `mapstructure:"api_key"`
}

type HttpConfig struct {
	Proxy *string `mapstructure:"proxy"`
}
