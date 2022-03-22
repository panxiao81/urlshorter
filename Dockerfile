FROM golang:alpine AS builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN apk add git && go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/urlshorter ./main/

FROM alpine:latest
WORKDIR /
COPY --from=builder /usr/local/bin/urlshorter /usr/local/bin/urlshorter
COPY main/config/config.yaml /config/config.yaml
EXPOSE 8080
CMD ["urlshorter"]