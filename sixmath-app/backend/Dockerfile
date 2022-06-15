# builder image
FROM golang:1.18.3-alpine as base
WORKDIR /builder
RUN apk add upx
ENV GO111MODULE=on CGO_ENABLED=0
COPY go.mod go.sum /builder/
RUN go mod download
COPY . /builder/
RUN go build \
  -ldflags "-s -w" \
  -o /builder/main /builder/main.go
RUN upx -9 /builder/main

# runner image
FROM alpine:latest
WORKDIR /app
COPY --from=base /builder/main main
COPY --from=base /builder/.env .env
EXPOSE 8000
CMD ["/app/main"]