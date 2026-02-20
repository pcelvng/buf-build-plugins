# buf-build-plugins

Custom public plugins for the Buf Schema Registry (BSR) ecosystem.

## Plugins

### service-no-forbidden-word

Enforces that Protobuf service names do not contain forbidden words (e.g. `Service`). Useful when requiring services to end with `API` instead of `Service`.

- [Documentation](plugins/service-no-forbidden-word/README.md)
- BSR: `buf.build/pcelvng/service-no-forbidden-word`

## Building & Publishing

### First-time publish

1. Ensure you're logged into Buf: `buf registry login`
2. Create the plugin on BSR (or use `--create` which creates it on first push)
3. Build and push:

```bash
make build
make push  # Uses BSR_ORG=pcelvng by default
```

4. In projects using the plugin, run `buf plugin update` to pin the version

### Subsequent publishes

```bash
make push
```

Each plugin lives under `plugins/<name>/` and can be built and published individually—see the plugin's README for details.
