all: build run
build:
	docker build -t apigateway .
run:
	docker run -d -p 8080:8080/tcp --name=apigateway apigateway:latest