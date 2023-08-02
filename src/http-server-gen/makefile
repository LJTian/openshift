
IMAGES_NAME ?="ljtian/http-server-gen"
VERSION ?="tmpe_version"

all: docker_build run

docker_build:
	docker build -f docker.der . -t ${IMAGES_NAME}:${VERSION}

docker_push:
	docker push ${IMAGES_NAME}:${VERSION}

run:
	docker run -P -d ${IMAGES_NAME}:${VERSION}