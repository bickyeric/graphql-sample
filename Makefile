db-up:
	go run cmd/migrator/main.go up

db-down:
	go run cmd/migrator/main.go down

db-reset:
	make db-down
	make db-up

db-seed:
	go run cmd/seeder/*.go

run:
	go run cmd/api/main.go