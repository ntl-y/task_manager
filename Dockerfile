FROM golang:1.23-alpine

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o task-manager cmd/main.go

CMD ["./task-manager"]