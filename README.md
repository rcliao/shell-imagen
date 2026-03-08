# shell-imagen

Image generation via Google Gemini API. CLI tool and Go library.

Part of the [Ghost in the Shell](https://github.com/rcliao?tab=repositories&q=shell-) ecosystem.

## Install

```bash
go install github.com/rcliao/shell-imagen/cmd/shell-imagen@latest
```

## Usage

```bash
# Generate and save to file
shell-imagen "a cute pikachu in watercolor style" --output pikachu.png

# Use specific model
shell-imagen "sunset over mountains" --model gemini-3.1-flash-image-preview --output sunset.png

# Output base64 to stdout (for piping)
shell-imagen "abstract art"
```

## Configuration

Set your Gemini API key:
```bash
export GEMINI_API_KEY="your-key-here"
```

Or pass via flag:
```bash
shell-imagen "prompt" --api-key "your-key"
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--api-key` | `$GEMINI_API_KEY` | Gemini API key |
| `--model` | `gemini-3.1-flash-image-preview` | Model to use |
| `--output` / `-o` | stdout (base64) | Output file path |
| `--timeout` | 2m | Request timeout |

## Library Usage

```go
import imagen "github.com/rcliao/shell-imagen"

gen, err := imagen.New(apiKey, "", 0)  // default model and timeout
imgBytes, err := imagen.Generate(ctx, gen, "a cute pikachu")
os.WriteFile("pikachu.png", imgBytes, 0644)
```

## Build

```bash
make build    # Build binary
make test     # Run tests
make vet      # Run go vet
```

## License

MIT
