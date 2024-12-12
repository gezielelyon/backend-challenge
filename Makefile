prepare:
	docker-compose up --build -d

down:
	docker-compose down

test:
	go test ./tests/... -v

swagger:
	swag init --dir ./cmd --output ./docs
