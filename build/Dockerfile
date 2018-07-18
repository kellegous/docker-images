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
		protobuf-compiler \
	&& ln -s /usr/bin/nodejs /usr/bin/node \
	&& npm install -g typescript \
	&& gem install rake sass \
	&& apt-get clean \
	&& curl https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz | tar xz -C /usr/local

ENV PATH="/usr/local/go/bin:${PATH}"
