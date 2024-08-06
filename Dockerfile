FROM golang:bookworm

WORKDIR /app

COPY ./src /app

CMD ["go", "run", "main.go"]