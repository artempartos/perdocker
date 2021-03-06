install: pull-image

docker-stop:
	sudo kill -QUIT `cat /var/run/docker.pid`

run: build run-perdocker

run-docker:
	sudo ./bin/docker -d
run-perdocker:
	./bin/perdocker

build:
	go build && mv perdocker ./bin/perdocker && chmod +x ./bin/perdocker

build-image: 
	docker build -rm -t="perdocker/universal:latest" ./images/universal/

build-images-lang: build-image-ruby build-image-nodejs build-image-go build-image-python build-image-c build-image-php

build-image-ruby:
	docker build -rm -t="perdocker/ruby:attach" ./images/ruby/
build-image-nodejs:
	docker build -rm -t="perdocker/nodejs:attach" ./images/nodejs/
build-image-go:
	docker build -rm -t="perdocker/go:attach" ./images/go/
build-image-python:
	docker build -rm -t="perdocker/python:attach" ./images/python/
build-image-c:
	docker build -rm -t="perdocker/c:attach" ./images/c/
build-image-php:
	docker build -rm -t="perdocker/php:attach" ./images/php/

pull-image: 
	docker pull perdocker/universal
pull-images-lang: pull-image-ruby pull-image-nodejs pull-image-go pull-image-python pull-image-c pull-image-php

pull-image-ruby:
	docker pull perdocker/ruby
pull-image-nodejs:
	docker pull perdocker/nodejs
pull-image-go:
	docker pull perdocker/go
pull-image-python:	
	docker pull perdocker/python
pull-image-c:	
	docker pull perdocker/c
pull-image-php:	
	docker pull perdocker/php

