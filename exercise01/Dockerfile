FROM golang:1.23.2
WORKDIR /app

COPY go.mod ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]