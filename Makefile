.PHONY: buid

buid:
	@go build -o bin/server cmd/server/main.go

piqo:
	@./bin/server