.PHONY: build

build-describer: clean
	CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -a -v -mod=mod -ldflags "-linkmode external -extldflags '-static' -s -w" -tags musl -o ./local/og-describer-github main.go

clean:
	rm -rf ./local/og-describer-github ./build/og-github-cli