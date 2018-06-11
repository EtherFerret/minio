FROM alpine:3.7
#FROM golang:1.10.1-alpine3.7

RUN \
     apk add --no-cache ca-certificates && \
     apk add --no-cache --virtual .build-deps curl openssl && \
     echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

COPY dockerscripts/docker-entrypoint.sh dockerscripts/healthcheck.sh /usr/bin/
COPY bin/minio /usr/bin/
ADD bin/librdkafka.tar.gz /
RUN ls -l /usr/lib

ENV MINIO_UPDATE off
ENV MINIO_ACCESS_KEY_FILE=access_key \
    MINIO_SECRET_KEY_FILE=secret_key

RUN \
     chmod +x /usr/bin/docker-entrypoint.sh && \
     chmod +x /usr/bin/healthcheck.sh


EXPOSE 9000

ENTRYPOINT ["/usr/bin/docker-entrypoint.sh"]

VOLUME ["/data"]

HEALTHCHECK --interval=30s --timeout=5s \
    CMD /usr/bin/healthcheck.sh

CMD ["minio"]
