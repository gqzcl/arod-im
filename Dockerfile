# Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style

FROM golang:1.18 AS builder

LABEL maintanier="gqzcl <gqzcl@qq.com>"

ENV GOPEOXY https://goproxy.cn,direct
COPY . /src/arod-im
WORKDIR /src/arod-im

# 安装必要的软件包和依赖包
# RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
#     sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
#     sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
USER root
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    make

# 安装 protoc
USER root
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.17.3/protoc-3.17.3-linux-x86_64.zip && \
    unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc $GOPATH/bin

# 初始化项目
USER root
RUN make init

# $GOPATH/bin添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 清理垃圾
# USER root
# RUN apt-get clean && \
#     rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
#     rm /var/log/lastlog /var/log/faillog

EXPOSE 8001
EXPOSE 8002
EXPOSE 8003
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003

# FROM debian:stable-slim

# COPY --from=builder /src/bin /app

# VOLUME /data/conf

# CMD ["./server", "-conf", "/data/conf"]
