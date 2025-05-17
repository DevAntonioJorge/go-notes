FROM golang:1.24.2-alpine

WORKDIR /app
COPY . /app

RUN go build -o main main.go


RUN go mod download
RUN go mod tidy

EXPOSE 8000

CMD ["./main"]
