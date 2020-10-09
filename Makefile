build:
	docker-compose build
up:
	docker-compose up -d
stop:
	docker-compose stop
down:
	docketr-compose down
work:
	docker exec -it mydog-api-app bash
db:
	docker exec -it mydog-api-mongo bash
