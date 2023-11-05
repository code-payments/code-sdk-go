FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go build -o app

EXPOSE 3000

CMD ["./app"]
