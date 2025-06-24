build:
	@go build -o urban-waddle cmd/main.go
	@echo "Build complete"

test:
	@go test -v ./...
	@echo "Test complete"

run:
	@./urban-waddle
	@echo "Server running on port 8080"

clean:
	@rm -f urban-waddle
	@echo "Clean complete"