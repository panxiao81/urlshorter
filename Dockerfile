FROM golang:bullseye AS builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/urlshorter ./main

FROM busybox:stable
WORKDIR /
COPY --from=builder /usr/local/bin/urlshorter /usr/local/bin
COPY main/config/config.yaml /config/config.yaml
EXPOSE 8080
CMD ['urlshorter']