build:
    go build -o api .

docker-build:
    docker build . -t api .

run: docker-build
    docker run -p 4242:4242 api

attatched: docker-build
    docker run -d -p 4242:4242 api
