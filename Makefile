docker-build:
	docker build . -t api

run: docker-build
	docker run -p 4242:4242 api

attached: docker-build
	docker run -d -p 4242:4242 api
