FROM golang:latest

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o main ./cmd

WORKDIR /dist

RUN cp /app/main .

CMD ["/dist/main"]