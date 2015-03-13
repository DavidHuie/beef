default: build

dep:
	godep save -r ./...

build: dep
	go build -o bin/beef github.com/DavidHuie/beef/cmd/beef

test: dep
	go test ./...
