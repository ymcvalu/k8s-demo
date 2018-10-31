TAG=kst-yapi:1.0.0

build:
	-docker rmi $(TAG)
	docker build -t $(TAG) .

push:
	docker tag $(TAG) registry.cn-qingdao.aliyuncs.com/yangmc/$(TAG)
	docker push registry.cn-qingdao.aliyuncs.com/yangmc/$(TAG)
