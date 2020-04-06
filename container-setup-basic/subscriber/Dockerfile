FROM golang:1.14

WORKDIR /go/src/app
RUN go get github.com/streadway/amqp
COPY . . 
CMD ["go", "run", "main.go"]

