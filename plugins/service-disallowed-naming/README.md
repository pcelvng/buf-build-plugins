# service-disallowed-naming

A Buf check plugin that enforces service names do not contain any of a user-defined list of disallowed words. Requires at least one configured disallowed word. You specify which words are not allowed based on your naming conventions.

## Rule: SERVICE_DISALLOWED_NAMING

Checks that service names do not contain any of the configured disallowed words. Matching is case-insensitive. If `forbidden_words` is not specified or empty, the rule does nothing.

## Usage

Add the plugin to your `buf.yaml` and configure the words you want to disallow:

```yaml
version: v2
modules:
  - path: proto
lint:
  use:
    - STANDARD
    - SERVICE_DISALLOWED_NAMING
plugins:
  - plugin: buf.build/pcelvng/service-disallowed-naming
    options:
      forbidden_words:
        - Service   # Disallow "Service" (e.g. when requiring "API" suffix)
        # Add any other words for your use case
```

Then run `buf plugin update` to pin the version, and `buf lint` to check your schemas.

## Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `forbidden_words` | `[]string` | none | List of words that must not appear in service names. Required for the rule to have any effect. |

## Example use cases

**Require "API" suffix instead of "Service":**
```yaml
forbidden_words:
  - Service
```

**Disallow multiple words:**
```yaml
forbidden_words:
  - Service
  - Endpoint
  - Handler
```

## Examples

With `forbidden_words: [Service]`:

```bash
# Passes (no disallowed words)
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
GOOS=wasip1 GOARCH=wasm go build -o service-disallowed-naming.wasm ./cmd/buf-plugin-service-disallowed-naming

# Push to BSR
buf plugin push buf.build/pcelvng/service-disallowed-naming \
  --binary=service-disallowed-naming.wasm \
  --create \
  --create-type=check \
  --create-visibility=public
```
