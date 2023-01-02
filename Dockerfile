FROM golang:latest

WORKDIR /src

COPY . /src

RUN go mod tidy

CMD ["go", "run", "main.go"]