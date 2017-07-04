# dependencies

godep save

# dockerize:

1) build bin first
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.4.2 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o myapp'

2) build image
docker build -t test-goapi .

3) run it
docker run maxhoffman/react-pg-test/goapi:latest

4) push to registry
docker tag test-goapi maxhoffman/react-pg-test/goapi:latest .
docker push maxhoffman/react-pg-test/goapi:latest

