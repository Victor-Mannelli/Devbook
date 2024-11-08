dev:
	# Run both tasks concurrently and keep them running
	# `-j 2` allows two jobs to run at once, for `build` and `tailwind`

	make -j 2 air tailwind

air:
	air main.go

tailwind:
	tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --watch

build:
	go build -o ./tmp/main .
