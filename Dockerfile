
FROM golang:latest AS builder

WORKDIR /build



ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod .
RUN go mod download

COPY . .
#RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s"  -o gokit cmd/main.go
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
EXPOSE 3003

