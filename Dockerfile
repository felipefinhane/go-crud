FROM golang:alpine3.13 AS base
RUN apk add --update --no-cache git ca-certificates perl openssh-client build-base
RUN go get github.com/codegangsta/gin
RUN go get golang.org/x/tools/cmd/godoc
RUN go get github.com/axw/gocov/gocov
RUN mkdir /db-test
WORKDIR /app
# Update go modules to latest version
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

FROM base AS dev
ENTRYPOINT sh -c "echo godoc at: http://$(show-ip|tr -d ' '):6060; (godoc -http 0.0.0.0:6060 > /dev/null &); echo application at: http://$(show-ip|tr -d ' '):3000; gin --appPort 9000 --bin build/gin-bin run main.go"

FROM base AS build
COPY . /app
RUN go build -o app .

FROM alpine:3.13 AS prod
WORKDIR /app
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/app /app/app
ENTRYPOINT sh -c "./app"
