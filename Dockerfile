FROM golang:1.19.3
WORKDIR /bot
COPY . /bot
CMD go run main.go