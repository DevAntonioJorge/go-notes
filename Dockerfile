FROM golang:1.24.2-alpine

WORKDIR /app
COPY . /app

RUN apk add --no-cache git

RUN go mod download
RUN go mod tidy

RUN go build -o main cmd/app/main.go
RUN chmod +x main
EXPOSE 8000

CMD ["/app/main"]
