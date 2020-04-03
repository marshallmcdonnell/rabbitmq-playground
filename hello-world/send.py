import pika

# Setup connection to broker on local machine
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

# Add recipient queue
channel.queue_declare('hello')

# Send "hello world" message
channel.basic_publish(exchange='', routing_key='hello', body='Hello World!')
print(" [x] Sent 'Hello World!'")

# Gently close connection to flush network buffers
connection.close()
