
lint:
	docker run -t --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v

create_model:
	@echo "Creating new model..."
	@go run -mod=mod entgo.io/ent/cmd/ent new ${model}

dev:
	@echo "Starting api in dev mode..."
	@env air
	@echo "api running!"

generate:
	@echo "Genrating...."
	@go generate ./...
	@echo "Genrating done!!"
