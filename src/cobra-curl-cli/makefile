
IMAGES_NAME ?="ljtian/tcurl"
VERSION ?="tmpe_version"

all: docker_build
	@echo "使用方式如下："
	@echo "docker run ${IMAGES_NAME}:${VERSION} -U=<httpServiceURL> -S=true -t=3 -D=<DBDNS>"

docker_build:
	docker build -f docker.der . -t ${IMAGES_NAME}:${VERSION}

docker_push:
	docker push ${IMAGES_NAME}:${VERSION}