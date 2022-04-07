NACOS_VERSION ?= 2.0.3

prepare:
	docker pull nacos/nacos-server:$(NACOS_VERSION)
	docker run --name nacos-quick -e MODE=standalone -p 8848:8848 -p 9848:9848 -d nacos/nacos-server:$(NACOS_VERSION)
    export serviceAddr=127.0.0.1
    export serverPort=8848
    export namespace=test
