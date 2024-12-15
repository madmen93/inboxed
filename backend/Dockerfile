FROM golang:1.21.12-alpine

RUN mkdir app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY ./cmd/indexer/index.json /app/cmd/indexer/index.json

RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]