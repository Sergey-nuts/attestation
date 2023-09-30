build-all: build-filter build-comments build-news build-gateway

build-filter:
	docker build -t filter Filter/.

build-comments:
	docker build -t comments Comments/.

build-news:
	docker build -t news News/.

build-gateway:
	docker build -t apigateway APIGateway/.


run-all: run-filter run-comments run-news run-gateway

run-filter:
	docker run -d -p 2020:2020/tcp --net=gateway --name=filter filter:latest

run-comments:
	docker run -d -p 2010:2010/tcp -e dbuser -e dbpass -e dbhost -e filter=filter --net=gateway --name=comments comments:latest

run-news:
	docker run -d -p 2000:2000/tcp -e dbuser -e dbpass -e dbhost --net=gateway --name=news news:latest
	
run-gateway:
	docker run -d -p 8080:8080/tcp -e news=news -e comments=comments --net=gateway --name=apigateway apigateway:latest

network:
	docker network create gateway