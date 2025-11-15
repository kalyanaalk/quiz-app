FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./bin/app

EXPOSE 8080

CMD ["./main"]
