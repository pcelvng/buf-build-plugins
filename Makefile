.PHONY: build push help

# BSR organization - update this for your BSR account
BSR_ORG ?= pcelvng
PLUGIN_NAME := service-no-forbidden-word

help:
	@echo "Targets:"
	@echo "  build    - Build the plugin (native + WASM)"
	@echo "  push     - Push plugin to BSR (default BSR_ORG=pcelvng)"
	@echo "  tidy     - Run go mod tidy in plugin directories"

build:
	cd $(PLUGIN_NAME) && go build -o buf-plugin-$(PLUGIN_NAME) ./cmd/buf-plugin-$(PLUGIN_NAME)
	cd $(PLUGIN_NAME) && GOOS=wasip1 GOARCH=wasm go build -o $(PLUGIN_NAME).wasm ./cmd/buf-plugin-$(PLUGIN_NAME)
	@echo "Built: $(PLUGIN_NAME)/buf-plugin-$(PLUGIN_NAME) (native) and $(PLUGIN_NAME)/$(PLUGIN_NAME).wasm (WASM)"

push: build
	cd $(PLUGIN_NAME) && buf plugin push buf.build/$(BSR_ORG)/$(PLUGIN_NAME) \
		--binary=$(PLUGIN_NAME).wasm \
		--create \
		--create-type=check \
		--create-visibility=public \
		--source-control-url=https://github.com/$(BSR_ORG)/buf-build-plugins

tidy:
	cd $(PLUGIN_NAME) && go mod tidy
