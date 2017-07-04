
# dockerize:

1) build bin first
docker run --rm -v "$PWD":/Go/src/github.com/react-pg-test -w /Go/src/github.com/react-pg-test iron/go:dev go build -o myapp

2) build image
docker build -t test-goapi .
docker tag test-goapi maxhoffman/react-pg-test/goapi:latest .
docker push maxhoffman/react-pg-test/goapi:latest

3) run it
docker run maxhoffman/react-pg-test/goapi:latest

