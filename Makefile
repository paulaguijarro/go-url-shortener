build:
	go build -o go-short main.go

run: build
	./go-short