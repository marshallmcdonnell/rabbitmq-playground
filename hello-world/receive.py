import pika

connection = pika.BlockingConnection(pika.ConnectionParamters('localhost'))
channel = connection.channel()

channel.queue_declare('hello')


