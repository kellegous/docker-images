FROM ubuntu:bionic

RUN apt-get update \
	&& apt-get install -y curl \
	&& curl -sL https://deb.nodesource.com/setup_14.x | bash - \
	&& apt-get update \
	&& apt-get install -y \
		git \
		ruby \
		ruby-dev \
		nodejs \
		build-essential \
		closure-compiler \
	&& npm install -g typescript \
	&& gem install rake sass \
	&& apt-get clean \
	&& curl -OL https://github.com/google/protobuf/releases/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip \	
  	&& unzip protoc-3.13.0-linux-x86_64.zip -d /usr/local \	
  	&& rm protoc-3.13.0-linux-x86_64.zip \
	&& curl https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz | tar xz -C /usr/local

ENV PATH="/usr/local/go/bin:${PATH}"
