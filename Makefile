dbup:
	docker run --rm --name pwebsite-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=pwebsite-db -d postgres

migrateup:
	migrate -database postgresql://root:secret@localhost:5432/pwebsite-db?sslmode=disable -path db/migrations up

migratedown:
	migrate -database postgresql://root:secret@localhost:5432/pwebsite-db?sslmode=disable -path db/migrations down