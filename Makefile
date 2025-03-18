# Database connection string (replace with actual credentials)
DB_URL=root:Maclocal12345@tcp(localhost:3306)/GoTwitter
MIGRATIONS_DIR=db/migrations

# Create a new migration file # gmake migrate-create name=create_users_table
migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Apply all pending migrations (Up) # gmake migrate-up
migrate-up:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up

# Rollback the last migration (Down) # gmake migrate-down
migrate-down:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down

# Rollback all migrations and reset database # gmake migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

# Show current migration status # gmake migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # gmake migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" redo

# Run specific migration version # gmake migrate-version version=20200101120000
migrate-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # gmake migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # gmake migrate-force version=20200101120000
migrate-force:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" force $(version)

# Print Goose help # gmake migrate-help
migrate-help:
	goose -h
