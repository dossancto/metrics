env/setup:
	cp .env.example .env

compose/up:
	docker-compose up -d

compose/down:
	docker-compose down

compose/restart:
	docker-compose restart

run:
	swag init
	go run .

build:
	go run .

db/update:
	cd migrations; npx prisma db push

db/reset:
	cd migrations; npx prisma migrate reset

migration/setup:
	cd migrations; npm i
	cp migrations/.env.example migrations/.env
