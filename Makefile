.PHONY: run
run: build
	./bin/main

.PHONY: build
build:
	go build -o bin/main .


.PHONY: compose/up
compose/up: compose/up/redis

.PHONY: compose/up/redis
compose/up/redis: compose/down
	docker-compose up -d redis

.PHONY: compose/down
compose/down:
	docker-compose down
