all: build run
build:
	docker build -t news .
run:
	docker run -d -p 2000:2000/tcp -e dbuser -e dbpass -e dbhost --name=news news:latest