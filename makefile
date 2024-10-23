# Include environment variables from .env file
include .env

# Variables
MIGRATE = migrate
MIGRATIONS_DIR = internal/database/migrations

# Default target
.PHONY: help
help: 
	@echo "Makefile for managing migrations and other commands"
	@echo "Usage:"
	@echo "  make migrations                              Show migration commands"
	@echo "  make migration name=<migration_name>         Create a new migration"
	@echo "  make migrate_up step=<number>                Apply migrations with a specified step"
	@echo "  make migrate_down step=<number>              Roll back migrations with a specified step"
	@echo "  make migrate_refresh                         Roll back all migrations and re-apply them"
	@echo "  make migrate_force version=<version>         Force a specific migration version"
	@echo ""
	@echo "Other Commands:"
	@echo "  make run-cli                                 Run the CLI application"

# Create a new migration
.PHONY: migration
migration:
	@if [ -z "$(name)" ]; then \
		echo "migration_name is undefined. Usage: make migration name=<migration_name>"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext=sql -dir=$(MIGRATIONS_DIR) -seq $(name)

# Apply migrations
.PHONY: migrate_up
migrate_up:
	@echo "Current step: $(step)"
	@if [ -z "$(step)" ]; then \
		$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up; \
	else \
		$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up $(step); \
	fi

# Roll back the last migration
.PHONY: migrate_down
migrate_down:
	@echo "Current step: $(step)"
	@if [ -z "$(step)" ]; then \
		$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down; \
	else \
		$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down $(step); \
	fi

# Refresh migrations
.PHONY: migrate_refresh
migrate_refresh:
	@echo "Rolling back all migrations..."
	$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down -all
	@echo "Re-applying all migrations..."
	$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

# Force a specific migration version
.PHONY: migrate_force
migrate_force:
ifndef version
	$(error version is undefined. Usage: make migrate_force version=<version>)
endif
	$(MIGRATE) -path=$(MIGRATIONS_DIR) -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $(version)

# Force a specific migration version
.PHONY: run-cli
run-cli:
	go run ./cmd/cli