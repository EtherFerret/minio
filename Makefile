lakeadmin:
	cd src/lakeadmin && python setup.py sdist	
	mv src/lakeadmin/dist/* docker/lakeadmin/packages/

gobuild:
	docker build -t gobuild docker/gobuild

kafka:
	docker run --rm  -v `pwd`/src:/go/src \
		-v `pwd`/docker/minio/bin:/go/bin gobuild \
		sh -c "cd /go/src/librdkafka && ./configure --prefix=/usr && make && make install && tar czvf /go/bin/librdkafka.tar.gz /usr/lib/librd*"

minio:
	docker run --rm  -v `pwd`/src:/go/src \
		-v `pwd`/docker/minio/bin:/go/bin gobuild \
		sh -c "cd /go/src/librdkafka && make install && go build -o /go/bin/minio github.com/minio/minio"

docker-minio:
	docker build -t ehualu.com/minio docker/minio


docker-admin:lakeadmin
	docker build -t ehualu.com/lakeadmin docker/lakeadmin


docker:docker-minio docker-admin
