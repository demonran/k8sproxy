package main

import (
	"github.com/spf13/cobra"
	"k8sproxy/internal/conn"
	"k8sproxy/pkg/options"
	"log"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "k8sproxy",
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return conn.Connect()
		},
	}
	var cfgFile string
	var baseURL string
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.k8sproxy/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&baseURL, "url", "u", "", "base url to get connection info")

	options.InitCfg(cfgFile, baseURL)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Exit: %s", err)
	}
}
