

run:
	APP_ENV=development GREEN_API_URL=https://api.green-api.com HTTP_SERVER_PORT=5912 go run cmd/main.go

templates:
	templ generate

build:
	go build -o bin/app ./cmd/main.go

prod:
	APP_ENV=development GREEN_API_URL=https://api.green-api.com HTTP_SERVER_PORT=5912 ./bin/app