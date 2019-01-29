TAG?=latest

build:
	docker build --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy -t stack360/faas-lambdroid:$(TAG) .

push:
	docker push stack360/faas-lambdroid:$(TAG)
