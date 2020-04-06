FROM golang:1.14

WORKDIR /go/src/app
RUN apt update && \
    apt install netcat -y && \
    rm -rf /var/lib/apt/lists/* && \
    go get github.com/streadway/amqp
COPY . . 
CMD ["bash", "./docker-start.sh"]

