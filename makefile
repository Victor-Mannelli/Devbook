air:
	air main.go

tailwind:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

build:
	go build -o ./tmp/main .
