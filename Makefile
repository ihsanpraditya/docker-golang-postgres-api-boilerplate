MIGRATION_PATH=./db/migrations
DB_URL=postgres://user:pass@localhost:5432/dbname?sslmode=disable

migration-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migration-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migration-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down
