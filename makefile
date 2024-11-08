dev:
	go run main.go

run css:
	npx tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch
