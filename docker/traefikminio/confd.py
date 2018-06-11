from kafka import KafkaConsumer
import json
import logging

def process_message(msg):
    try:
        ev = json.loads(msg.value)
        if not ev.get('action') in ['user.add', 'user.remove']:
            return

    except Exception as e:
        logging.exception(e)
        continue

def recieve_user_event():
    consumer = KafaConsumer(bootstrap_servers='kafka:1234', topic='user-events', group_id='traefik')
    for msg in consumer:
        process_message(msg)

if __name__ == '__main__':
    try:
        recieve_user_event()
    except Exception as e:
        logging.exception(e)
        continue
