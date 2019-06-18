VERSION := "test-run"

build:
	mkdir -p build
	go build -ldflags "-s -w -X main.version=$(VERSION)" -o build/mxp-server command/main.go

clean:
	rm -rf build
