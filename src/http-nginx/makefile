
dockerRun:
	docker run -d --name=https-basic -p 443:443 \
	       	-v /mnt/data/shareDir_PC/git/go-test/openshift/http-nginx/conf.d:/etc/nginx/conf.d \
		-v /mnt/data/shareDir_PC/git/go-test/openshift/http-nginx/tls:/tls/ssl/  \
		nginx
dockerClean:
	docker stop https-basic
	docker rm https-basic
