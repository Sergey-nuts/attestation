all: build run
build:
	docker build -t comments .
run:
	docker run -d -p 2010:2010/tcp -e dbuser -e dbpass -e dbhost --name=comments comments:latest