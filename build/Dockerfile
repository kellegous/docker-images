FROM ubuntu:xenial

RUN apt-get update \
  && apt-get install -y golang git ruby nodejs npm build-essential closure-compiler curl \
  && ln -s /usr/bin/nodejs /usr/bin/node \
  && curl -OL https://github.com/google/protobuf/releases/download/v3.1.0/protoc-3.1.0-linux-x86_64.zip \
  && unzip protoc-3.1.0-linux-x86_64.zip -d /usr/local \
  && rm protoc-3.1.0-linux-x86_64.zip \
  && npm install -g typescript \
  && gem install rake sass \
  && apt-get clean
