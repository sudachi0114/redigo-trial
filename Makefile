.PHONY: run
run: build
	./bin/main

.PHONY: build
build:
	go build -o bin/main .


.PHONY: compose/up
compose/up: compose/up/redis
	docker-compose up -d app

.PHONY: compose/up/sh
compose/up/sh: compose/up
	docker-compose exec app sh

.PHONY: compose/up/redis
compose/up/redis: compose/down
	docker-compose up -d redis

.PHONY: compose/down
compose/down:
	docker-compose down

.PHONY: compose/build
compose/build:
	docker-compose build app
