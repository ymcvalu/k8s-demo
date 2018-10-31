TAG=k8s-test:1.0.0

build:
	-docker rmi $(TAG)
	docker build -t $(TAG) .
run:
	docker run --rm -p 8088:80 --name k8s-test -d $(TAG)
stop:
	docker stop k8s-test
