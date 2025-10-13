package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/stefan79/openai-driver/pkg/config"
	"github.com/stefan79/openai-driver/pkg/output"

	"github.com/stefan79/openai-driver/cmd/responses"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CtxKey int

const CfgKey CtxKey = 1
const OutputterKey CtxKey = 2

var (
	verboseLevel int
	cfgFile      string

	rootCmd = &cobra.Command{
		Use:   "oai",
		Short: "OpenAI CLI",
		Long:  `A CLI for interacting with the OpenAI API`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'oai --help' for more information")
		},
		PersistentPreRunE: initCommand,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oai.yaml)")
	rootCmd.PersistentFlags().IntVar(&verboseLevel, "verbose", 0, "output format (default is json)")
	rootCmd.AddCommand(responses.ResponsesCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func initCommand(cmd *cobra.Command, _ []string) error {
	outputter, err := initOutputter(cmd)
	if err != nil {
		return err
	}
	config, err := initConfig(cmd)
	if err != nil {
		return err
	}
	ctx := context.WithValue(cmd.Context(), CfgKey, config)
	ctx = context.WithValue(ctx, OutputterKey, outputter)
	cmd.SetContext(ctx)
	return nil
}

func initOutputter(cmd *cobra.Command) (output.Outputter, error) {
	if verboseLevel < -1 || verboseLevel > 3 {
		return nil, fmt.Errorf("invalid verbose level: %d. Needs to be between -1 and 3", verboseLevel)
	}
	outputter := output.NewDefaultOutputter()
	outputter.SetLevel(verboseLevel)
	return outputter, nil
}

func initConfig(cmd *cobra.Command) (*config.Config, error) {
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
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, nil
}
