# xzzpig/alidnsingressupdater
#build stage
FROM golang:1.15.7-alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
 && apk add --no-cache git
WORKDIR /go/src/app
COPY . .
ENV GOPROXY=https://goproxy.io,direct
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
 && apk --no-cache add ca-certificates
COPY --from=builder /go/bin/alidns-ingress-updater /alidns-ingress-updater
ENTRYPOINT ./alidns-ingress-updater
