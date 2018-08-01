FROM ubuntu:xenial

RUN apt-get update \
	&& apt-get install -y \
		git \
		ruby \
		ruby-dev \
		nodejs \
		npm \
		build-essential \
		closure-compiler \
		curl \
	&& ln -s /usr/bin/nodejs /usr/bin/node \
	&& npm install -g typescript \
	&& gem install rake sass \
	&& apt-get clean \
	&& curl -OL https://github.com/google/protobuf/releases/download/v3.1.0/protoc-3.1.0-linux-x86_64.zip \	
  	&& unzip protoc-3.1.0-linux-x86_64.zip -d /usr/local \	
  	&& rm protoc-3.1.0-linux-x86_64.zip \
	&& curl https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz | tar xz -C /usr/local

ENV PATH="/usr/local/go/bin:${PATH}"
