# service-no-forbidden-word

A Buf check plugin that enforces service names do not contain forbidden words (e.g. `Service`). Useful when you require services to end with `API` instead of `Service` and want to prevent names like `TxtServiceAPI`.

## Rule: SERVICE_NO_FORBIDDEN_WORD

Checks that Protobuf service names do not contain any of the configured forbidden words. Matching is case-insensitive.

## Usage

Add the plugin to your `buf.yaml`:

```yaml
version: v2
modules:
  - path: proto
lint:
  use:
    - STANDARD
    - SERVICE_NO_FORBIDDEN_WORD
  service_suffix: API  # Require API suffix
plugins:
  - plugin: buf.build/pcelvng/service-no-forbidden-word
    options:
      forbidden_words:
        - Service
```

Then run `buf plugin update` to pin the version, and `buf lint` to check your schemas.

## Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `forbidden_words` | `[]string` | `["Service"]` | List of words that must not appear in service names |

## Examples

```bash
# Passes (no forbidden words)
service TxtAPI {}

# Fails (contains "Service")
service TxtServiceAPI {}
```

## Publishing to BSR

From the repo root:

```bash
make build
make push
```

Or from this directory:

```bash
# Build WASM binary (required for BSR)
GOOS=wasip1 GOARCH=wasm go build -o service-no-forbidden-word.wasm ./cmd/buf-plugin-service-no-forbidden-word

# Push to BSR
buf plugin push buf.build/pcelvng/service-no-forbidden-word \
  --binary=service-no-forbidden-word.wasm \
  --create \
  --create-type=check \
  --create-visibility=public
```
