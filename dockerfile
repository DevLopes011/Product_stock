FROM golang:1.23

WORKDIR /go/src/app

COPY . .

RUN go build -o main ./cmd

EXPOSE 8000

CMD ["./main"]
