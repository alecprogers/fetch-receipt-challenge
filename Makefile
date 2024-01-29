.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/process-receipt-handler src/process-receipt-handler/*.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/process-receipt src/process-receipt/*.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
