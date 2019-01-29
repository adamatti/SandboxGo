FROM golang:1.11.5-stretch as builder

RUN mkdir /app
WORKDIR /app

# Dependencies
COPY go.mod .
RUN go mod download

# Build
COPY main.go .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" -s -w'

#############################################################
FROM alpine:latest as runner

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/crud .

EXPOSE 8000
CMD /app/crud