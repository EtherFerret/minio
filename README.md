
# How to build minio

git clone http:///10.2.174.241:8088/ehualu/datalakev2

export GOROOT=/path/to/go/
export PATH=$PATH:$GOROOT/bin
export GOPATH=/path/to/datalakev2

cd datalakev2/src/github.com/minio/minio
make

