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
	options.SetOptions(rootCmd, rootCmd.Flags(), options.GetOption(), options.OptionFlags())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Exit: %s", err)
	}
}
