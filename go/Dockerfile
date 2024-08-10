FROM golang:bookworm

WORKDIR /go_server

COPY . /go_server

RUN go mod tidy

RUN go build -o main

CMD ["./main"]