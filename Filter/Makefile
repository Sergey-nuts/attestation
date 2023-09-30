all: build run
build:
	docker build -t filter .
run:
	docker run -d -p 2020:2020/tcp -e dbuser -e dbpass -e dbhost --name=filter filter:latest