run-local: wire swag build
	./bin/app
build:
	go build -o bin/app cmd/app/main.go
swag:
	swag init --md ./docs --parseInternal --parseDependency --parseDepth 2 -g cmd/app/main.go
wire:
	google-wire ./internal/di