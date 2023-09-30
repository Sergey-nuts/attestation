FROM golang:1.21-alpine3.18 AS builder
RUN mkdir -p /app/src
WORKDIR /app/src
COPY . ./
RUN go build -o /app/News cmd/news/News.go

FROM alpine:3.18
WORKDIR /opt
EXPOSE 2000
COPY --from=builder /app/src/config.json /opt/config.json
COPY --from=builder /app/News /opt/news
CMD ["/opt/news"]