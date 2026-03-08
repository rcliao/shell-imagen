package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	shellimagen "github.com/rcliao/shell-imagen"
	"github.com/spf13/cobra"
)

func main() {
	var (
		apiKey  string
		model   string
		output  string
		timeout time.Duration
	)

	root := &cobra.Command{
		Use:   "shell-imagen [prompt]",
		Short: "Generate images using Google Gemini",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			prompt := strings.Join(args, " ")

			if apiKey == "" {
				apiKey = os.Getenv("GEMINI_API_KEY")
			}
			if apiKey == "" {
				return fmt.Errorf("API key required: set --api-key or GEMINI_API_KEY env")
			}

			g, err := shellimagen.New(apiKey, model, timeout)
			if err != nil {
				return fmt.Errorf("init generator: %w", err)
			}

			data, err := g.Generate(cmd.Context(), prompt)
			if err != nil {
				return fmt.Errorf("generate: %w", err)
			}

			if output != "" {
				if err := os.WriteFile(output, data, 0o644); err != nil {
					return fmt.Errorf("write file: %w", err)
				}
				fmt.Fprintf(os.Stderr, "wrote %s (%d bytes)\n", output, len(data))
				return nil
			}

			// No output file: write base64 to stdout
			fmt.Println(base64.StdEncoding.EncodeToString(data))
			return nil
		},
	}

	root.Flags().StringVar(&apiKey, "api-key", "", "Google Gemini API key (or GEMINI_API_KEY env)")
	root.Flags().StringVar(&model, "model", "", "model name (default: gemini-3.1-flash-image-preview)")
	root.Flags().StringVarP(&output, "output", "o", "", "output file path (default: base64 to stdout)")
	root.Flags().DurationVar(&timeout, "timeout", 0, "generation timeout (default: 2m)")

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
