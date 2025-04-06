FROM golang:1.24.1

RUN mkdir /app

ADD . /app

WORKDIR /app

ADD .env /app

RUN go build -o main cmd/main.go

CMD ["/app/main"]