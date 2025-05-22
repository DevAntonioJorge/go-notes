FROM golang:1.24.2-alpine

WORKDIR /app
COPY . /app

RUN go build -mod=vendor -o main cmd/app/main.go
RUN chmod +x main
EXPOSE 8000

CMD ["/app/main"]
