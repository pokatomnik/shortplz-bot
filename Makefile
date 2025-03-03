BINARY_NAME=shortplz

BUILD_DIR=build

PLATFORM?=$(shell go env GOOS)

GO=go

LDFLAGS=-ldflags "-s -w"

TRIMPATH=-trimpath

all: build

build:
	@echo "Building for $(PLATFORM)..."
	@mkdir -p $(BUILD_DIR)
	@if [ "$(PLATFORM)" = "win32" ]; then \
		GOOS=windows GOARCH=amd64 $(GO) build ${TRIMPATH} $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME).exe .; \
	elif [ "$(PLATFORM)" = "linux" ]; then \
		GOOS=linux GOARCH=amd64 $(GO) build ${TRIMPATH} $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .; \
	elif [ "$(PLATFORM)" = "darwin" ]; then \
		GOOS=darwin GOARCH=arm64 $(GO) build ${TRIMPATH} $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .; \
	else \
		echo "Unsupported platform: $(PLATFORM)"; \
		exit 1; \
	fi

clean:
	rm -rf $(BUILD_DIR)

.PHONY: all build clean