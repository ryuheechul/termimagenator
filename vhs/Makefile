.PHONY: vhs
vhs: build images
	PATH=$(PWD):$(PATH) vhs demo.tape

.PHONY: images
images:
	cat ./images.txt | xargs -I _ docker pull _

.PHONY: build
build:
	go build ../cmd/tmi

.PHONY: rmi-all
rmi-all:
	docker rmi -f $(shell docker images -aq)
