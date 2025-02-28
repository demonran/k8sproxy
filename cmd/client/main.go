package main

import (
	"github.com/spf13/cobra"
	"k8sproxy/internal/conn"
	"k8sproxy/pkg/options"
	"log"
)

func main() {
	log.Printf("start")
	rootCmd := &cobra.Command{
		Use:     "k8sproxy",
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return conn.Connect()
		},
	}

	options.Init()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Exit: %s", err)
	}
}
