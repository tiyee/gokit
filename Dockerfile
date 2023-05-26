
FROM golang:latest AS builder

WORKDIR /build



ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct


COPY . .
RUN make build
FROM alpine:latest AS final

WORKDIR /app
COPY --from=builder /build/bin/gokit /app/

RUN echo "https://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories \
    && apk update  \
    && apk upgrade  \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata

EXPOSE 4718

