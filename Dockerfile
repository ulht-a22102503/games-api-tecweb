FROM golang:1.23 as builder
ADD ./app /app
WORKDIR /app
RUN go build -o /go/bin/api /app/*.go

FROM ubuntu:latest
EXPOSE 50007
RUN mkdir -p /go/bin
COPY --from=builder /go/bin/api /go/bin
WORKDIR /go/bin
ENTRYPOINT "./api"