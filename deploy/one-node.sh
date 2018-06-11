#docker pull spotify/kafka
#docker pull mongo

# Make sure ak, sk, ceph_url is corret
ADMIN_AK=9CRPC6W8EHIOJGOQ93Q6
ADMIN_SK=EuOllV6rcFw7rw3aYAN4K1VPVBvO4e2TsG58AOkW
CEPH_URL=http://10.88.88.136:8080


# Create network
docker network create datalake -d overlay

# Create public services
#docker service create --name kafka --network datalake wurstmeister/kafka 
docker service create  --name mongo --network datalake mongo
docker service create  --name kafka --network datalake spotify/kafka



# admin service
docker service create --name lakeadmin --network datalake \
    -e ADMIN_AK=$ADMIN_AK \
    -e ADMIN_SK=$ADMIN_SK \
    -e CEPH_RGW_URL=$CEPH_URL \
    -e KAFKA_EP=kafka:9092 \
    -p 8080:8080 \
    ehualu.com/lakeadmin

# minio service
docker service create --name minio --network datalake \
    -e MINIO_ACCESS_KEY=$ADMIN_AK \
    -e MINIO_SECRET_KEY=$ADMIN_SK \
    -e LAKEADMIN_EP=http://lakeadmin:8080 \
    -e KAFKA_EP=kafka:9092 \
    -p 9000:9000 \
    ehualu.com/minio gateway s3 $CEPH_URL
