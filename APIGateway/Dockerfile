FROM golang:1.21-alpine3.18 AS builder
RUN mkdir -p /app/src
WORKDIR /app/src
COPY . ./
RUN go build -o /app/APIGateway cmd/apigateway/APIGateway.go

FROM alpine:3.18
WORKDIR /opt
EXPOSE 8080
COPY --from=builder /app/APIGateway /opt/apigateway
CMD ["/opt/apigateway"]