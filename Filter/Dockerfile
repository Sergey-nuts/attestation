FROM golang:1.21-alpine3.18 AS builder
RUN mkdir -p /app/src
WORKDIR /app/src
COPY . ./
RUN go build -o /app/Filter cmd/filter/Filter.go

FROM alpine:3.18
EXPOSE 2000
COPY --from=builder /app/Filter /opt/Filter
CMD ["/opt/Filter"]