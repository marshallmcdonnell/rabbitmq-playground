package main

import (
    "log"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

// AMQP client API: https://godoc.org/github.com/streadway/amqp

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-server:5672/")
    failOnError(err, "Error connection to the broker")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    exchangeName := "user_updates"
    bindingKey   := "user.profile.*"

    // Create the exchange if it doesn't already exist
    err = ch.ExchangeDeclare(
        exchangeName,   // name
        "topic",        // kind
        true,           // durable
        false,          // autoDelete
        false,          // internal
        false,          // noWait
        nil,            // args
    )
    failOnError(err, "Error creating the exchange")

    // Create the queue if it doesn't already exist
    queue, err := ch.QueueDeclare(
        "",     // name - empty means a random, unique name assigned
        true,   // durable
        false,  // autoDelete
        false,  // exclusive
        false,  // noWait
        nil,    // args
    )
    failOnError(err, "Error creating the queue")

    // Bind the queue to the exchange
    err = ch.QueueBind(
        queue.Name,     // name
        bindingKey,     // key
        exchangeName,   // exchange
        false,          // noWait
        nil,            // args
    )

    // Subscribe to the queue
    msgs, err := ch.Consume(
        queue.Name,     // name
        "",             // consumer - empty means a random, unique ID assigned
        false,          // autoAck
        false,          // exclusive
        false,          // noLocal
        false,          // noWait
        nil,            // args
    )
    failOnError(err, "Failed to register as a consumer")

    // Initialize the forever channel
    forever := make(chan bool)

    // Consumer goroutine function
    go func() {
        for d := range msgs {
            log.Printf("Received message: %s", d.Body)
            d.Ack(
                false,  // multiple - acknowledge for multiple deliveris
            )
        }
    }()

    log.Printf("Service listening for events...")

    // Block until forever receives a value, which will never happen.
    <-forever
}
