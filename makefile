dev:
	go run main.go

run css:
	npx tailwindcss -i ./src/css/input.css -o ./src/dist/output.css --watch
