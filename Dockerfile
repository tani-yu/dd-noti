FROM golang:1.17.7-alpine AS builder
WORKDIR /go/src/github.com/tani-yu/dd-noti
COPY ./ /go/src/github.com/tani-yu/dd-noti
RUN apk update \
    && apk add --no-cache git openssh libc-dev \
    && cd /go/src/github.com/tani-yu/dd-noti \
    && go mod tidy \
    && go build -o docker-dd-noti main.go

FROM alpine
COPY --from=builder /go/src/github.com/tani-yu/dd-noti/docker-dd-noti /usr/local/bin/docker-dd-noti
RUN apk update \
    && apk add ca-certificates --no-cache
CMD ["/usr/local/bin/docker-dd-noti"]