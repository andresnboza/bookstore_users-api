all: clean build run

clean:
	rm -f *.mod

build:
	export GO111MODULE=on
	go mod init
	go mod tidy
	go build

run:
	go run main.go