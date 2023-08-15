dbup:
	docker run --rm --name pwebsite-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=pwebsite-db -d postgres

migrateup:
	migrate -database postgresql://root:secret@localhost:5432/pwebsite-db?sslmode=disable -path build/db/migrations up

migratedown:
	migrate -database postgresql://root:secret@localhost:5432/pwebsite-db?sslmode=disable -path build/db/migrations down

fmt:
	go fmt ./...

run:
	go run cmd/pwebsite-backend.go

clean: 
	rm -rf internal/generated/*

gen: clean
	sqlc generate
