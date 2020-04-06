#!/bin/sh

set -x

until nc -vz $RABBITMQ_HOSTNAME $RABBITMQ_PORT; do
  >&2 echo "RabbitMQ is unavailble - sleeping"
  sleep 1
done
echo "RabbitMQ is ready, starting Python publisher"

go run main.go
