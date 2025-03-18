# Database connection string (replace with actual credentials)
DB_URL=user:password@tcp(localhost:3306)/dbname
MIGRATIONS_DIR=db/migrations

# Create a new migration file
migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Apply all pending migrations (Up)
migrate-up:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up

# Rollback the last migration (Down)
migrate-down:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down

# Rollback all migrations and reset database
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

# Show current migration status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

# Redo last migration (Down then Up)
migrate-redo:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" redo

# Run specific migration version
migrate-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version
migrate-down-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version
migrate-force:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" force $(version)

# Print Goose help
migrate-help:
	goose -h
