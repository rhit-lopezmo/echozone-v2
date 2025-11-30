DEFAULT_GOAL: build-tui

.PHONY: tidy run-tui build-tui

tidy:
	go mod tidy

run-tui: tidy
	go run ./cmd/echozone-tui

build-tui: tidy
	go build -o ez-tui ./cmd/echozone-tui
