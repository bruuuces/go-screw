.PHONY: all build darwin-amd64 darwin-arm64 windows-amd64 linux-amd64 linux-arm64 clean help

APP=gtool

help:
		@echo "Usage: make [target]"
		@echo ""
		@echo "Targets:"
		@echo "  deploy          Deploy the binary to GitHub"

deploy: build
		@echo "Deploying $(APP)..."
		@git tag -a v$(VERSION) -m "Release v$(VERSION)"
		@git push origin v$(VERSION)