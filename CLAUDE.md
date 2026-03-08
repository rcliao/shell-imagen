# shell-imagen

Image generation via Google Gemini API.

## Architecture

- `cmd/shell-imagen/main.go` — Cobra CLI: prompt arg, --api-key, --model, --output, --timeout flags
- `internal/imagen/imagen.go` — Generator struct wrapping `google.golang.org/genai` client
- `internal/imagen/imagen_test.go` — Tests
- `imagen.go` — Public API: `Generator` type alias, `New()`, `Generate()`

## Build & Test

```bash
make build    # Build binary
make test     # Run tests
make vet      # Run go vet
```

## Key Patterns

- `New(apiKey, model, timeout)` creates a Generator (defaults: model `gemini-3.1-flash-image-preview`, timeout 2min)
- `Generate(ctx, generator, prompt)` returns raw image bytes (`[]byte`)
- Uses `google.golang.org/genai` SDK
- CLI writes to file (--output) or base64 to stdout
