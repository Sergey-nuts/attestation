FROM golang:1.21-alpine3.18 AS builder
RUN mkdir -p /app/src
WORKDIR /app/src
COPY . ./
RUN go build -o /app/Comments cmd/comments/Comments.go

FROM alpine:3.18
EXPOSE 2010
COPY --from=builder /app/Comments /opt/Comments
CMD ["/opt/Comments"]