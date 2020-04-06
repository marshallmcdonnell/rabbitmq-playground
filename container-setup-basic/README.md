## To Run

Spin up contianers (RabbitMQ server, Python publisher, and Go subscriber)
```
docker-compose up
```

Then send commands to the `/users/<id>` endpoint
```
curl -d "full_name=me" -X POST http://<server>:3000/users/30
```
