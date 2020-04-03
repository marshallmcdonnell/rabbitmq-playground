import pika
import json

def emit_user_profile_update(user_id, new_data):
    params = pika.ConnectionParameters(host='rabbitmq-server')
    connection = pika.BlockingConnection(params)
    channel = connection.channel()

    exchange_name = "user_updates"
    routing_key = "user.profile.update"

    channel.exchange_declare(exchange=exchange_name, exchange_type='topic', durable=True)
    
    new_data['id'] = user_id

    channel.basic_publish(
        exchange=exchange_name,
        routing_key=routing_key,
        body=json.dumps(new_data),
        properties=pika.BasicProperties(delivery_mode=2))

    msg = "{routing_key} send to exchange {exchange_name} with data: {data}"
    print(msg.format(
        routing_key=routing_key,
        exchange_name=exchange_name,
        data=new_data))
    connection.close()
