package cmd

import (
	"fmt"
	"strings"

	"github.com/stefan79/openai-driver/pkg/config"

	"github.com/stefan79/openai-driver/cmd/files"
	"github.com/stefan79/openai-driver/cmd/responses"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "oai",
		Short: "OpenAI CLI",
		Long:  `A CLI for interacting with the OpenAI API`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'oai --help' for more information")
		},
		PersistentPreRunE: initConfig,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oai.yaml)")

	rootCmd.AddCommand(responses.ResponsesCmd)
	rootCmd.AddCommand(files.FilesCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig(cmd *cobra.Command, _ []string) error {
	v := viper.New()
	cfg := config.Config{}

	// 1) Defaults to Set
	cfg.OpenAI.BaseUrl = "https://api.openai.com"
	cfg.OpenAI.Model = "gpt-5"
	cfg.Console.Output = "json"

	// 2) ConfigFiles
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
		if err := v.ReadInConfig(); err != nil {
			return fmt.Errorf("read config: %w", err)
		}
	} else {
		v.SetConfigName("oai")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME/.oai")
		_ = v.ReadInConfig() // ignore not-found
	}

	// 3) Env overides
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	if err := v.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("unmarshal config: %w", err)
	}

	config.SetContext(cmd, &cfg)

	return nil
}
