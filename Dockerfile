FROM golang:bookworm

WORKDIR /app

COPY ./src /app

RUN go mod tidy

RUN go build -o main

CMD ["./main"]
