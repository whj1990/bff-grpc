.PHONY: init
init:
	go install google.golang.org/grpc@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: wire
wire:
	wire

.PHONY: swag
swag:
	swag init  --parseDependency