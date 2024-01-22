
install:
	npm install
	go install github.com/joho/godotenv/cmd/godotenv@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/cosmtrek/air@latest

run:
	go run main.go

deploy:
	fly deploy

dev:
	WATCH=true godotenv -f .env air

air:
	make templ
	make css_dev
	go build -o ./tmp/main main.go
	clear

css:
	npx tailwindcss -i styles.css -o ./static/styles.css -m

css_dev:
	npx tailwindcss -i styles.css -o ./static/styles.css

templ:
	templ generate

fmt:
	templ fmt .
	go fmt .
