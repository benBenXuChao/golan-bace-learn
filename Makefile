TARGET=api
DOCKER_TAG= temp  #  $(if $(DOCKERTAG),$(DOCKERTAG),latest)
DOCKER_TARGET= hub.deepin.com/wuhan_web_service_certify/api:$(DOCKER_TAG)

CURDIR=$(shell pwd)

build:
	go build -mod vendor -o $(TARGET) $(CURDIR)/cmd

docker-clear:
	docker rmi $(docker images|grep none|awk '{print $3 }')

docker:
	docker build --build-arg TARGET=$(TARGET) \
		-f Dockerfile \
		-t $(DOCKER_TARGET) .

docker-release:
	docker push $(DOCKER_TARGET)