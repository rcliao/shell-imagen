// Package shellimagen provides image generation using Google's Gemini API.
//
// This is the public API surface. The implementation lives in internal/imagen.
package shellimagen

import (
	"context"
	"time"

	"github.com/rcliao/shell-imagen/internal/imagen"
)

// Generator wraps the Gemini API for image generation.
type Generator = imagen.Generator

// New creates a new Generator. If model is empty, it defaults to
// "gemini-3.1-flash-image-preview". If timeout is zero, it defaults to 2 minutes.
func New(apiKey, model string, timeout time.Duration) (*Generator, error) {
	return imagen.New(apiKey, model, timeout)
}

// Generate produces an image from the given prompt and returns the raw image bytes.
func Generate(ctx context.Context, g *Generator, prompt string) ([]byte, error) {
	return g.Generate(ctx, prompt)
}
