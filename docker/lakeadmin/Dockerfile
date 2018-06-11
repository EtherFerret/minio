FROM daocloud.io/python:3.5-slim


RUN pip install requests boto3\
    docker rgwadmin bottle==0.12.13 cherrypy==8.9.1 \
    wsgi-request-logger prometheus_client retrying \
    gevent bottle-websocket kafka

COPY packages/* /packages/

RUN find packages -exec pip install {} \;

EXPOSE 8080

CMD ["lakeadmind"]
