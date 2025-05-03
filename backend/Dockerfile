FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ascii-server ./cmd/

EXPOSE 6060

CMD ["./ascii-server"]