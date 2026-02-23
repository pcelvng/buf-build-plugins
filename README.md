# buf-build-plugins

Custom public plugins for the Buf Schema Registry (BSR) ecosystem.

## Plugins

### service-disallowed-naming

Enforces that service names do not contain any of a user-defined list of disallowed words. Requires at least one configured disallowed word. Configure `forbidden_words` for your naming conventions (e.g. disallow "Service" when requiring "API" suffix).

- [Documentation](plugins/service-disallowed-naming/README.md)
- BSR: `buf.build/pcelvng/service-disallowed-naming` (pin a version with `:v0.1.0`)

## Building & Publishing

### First-time publish

1. Log in to BSR: `make login` (or `buf registry login`)
2. Build and push:

```bash
make build
make push
```

Uses `BSR_ORG=pcelvng` by default. In projects using the plugin, run `buf plugin update` to pin the version.

### Versioned publishes

After tagging a release in git (e.g. `v0.1.0`), push the plugin with that version label:

```bash
make push-version
```

The version is read from the `VERSION` variable in the Makefile (default: `v0.1.0`). Update it before each release, or override: `make push-version VERSION=v0.2.0`.

### Subsequent publishes (no version label)

```bash
make push
```

Each plugin lives under `plugins/<name>/` and can be built and published individually—see the plugin's README for details.
