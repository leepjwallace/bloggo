APP=bloggo
GOOS?=linux
GOARCH?=amd64

.PHONY: build
build:
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${APP}_${GOOS}_${GOARCH}

.PHONY: build-all
build-all:
	GOOS=linux GOARCH=amd64 go build -o ${APP}_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o ${APP}_darwin_amd64
	GOOS=linux GOARCH=arm64 go build -o ${APP}_linux_arm64
	GOOS=darwin GOARCH=arm64 go build -o ${APP}_darwin_arm64
