package constants

/*
 messages
*/
const NO_TARGET_SUPPLIED = "Must supply a target golang project "
const NOT_VALID_TARGET = "Cannot use provided directory as build context "
const NOT_VALID_GO2DOCKER_FILE = "Not valid go2docker.yml file "

/*
 app constants
*/
const GO2DOCKER_FILE = "go2docker.yml"
const CTX_SUFFIX = ".dkrctx"
const DOCKERFILE = "Dockerfile"
const LATEST_GO_IMAGE_VERSION = "latest"

const IMAGE_VERSION = "__IMAGE_VERSION__"
const APP_NAME = "__APP_NAME__"
const EXEC_NAME = "__EXEC_NAME__"

const DOCKERFILE_TEMPLATE = `
FROM golang:__IMAGE_VERSION__ AS builder

RUN mkdir -p /go/src/__APP_NAME__
COPY . /go/src/__APP_NAME__
WORKDIR /go/src/__APP_NAME__
RUN go get ./...
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o __EXEC_NAME__

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/__APP_NAME__/__EXEC_NAME__ /app/__EXEC_NAME__
CMD ["/app/__EXEC_NAME__"]
`
