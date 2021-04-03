.PHONY: run
run: build
	./bin/main

.PHONY: build
build:
	go build -o bin/main .


.PHONY: container/redis
container/redis:
	docker run -d -p 6379:6379 --name myredis redis

.PHONY: container/redis/console
container/redis/console:
	docker exec -it myredis bash

.PHONY: container/redis/rm
container/redis/rm:
	docker stop myredis && docker rm myredis
