# Makefile Commands

This document provides an overview of the commands available in the Makefile for managing database migrations and running the CLI application.

## Environment Variables
The Makefile includes environment variables from the `.env` file.

## Variables
- **MIGRATE**: Command used for managing migrations.
- **MIGRATIONS_DIR**: Directory where migration files are stored (`internal/database/migrations`).

## Default Target
### `make help`
Displays the available commands and their usage.

#### Output:
- **Usage**:
  - `make migrations`: Show migration commands.
  - `make migration name=<migration_name>`: Create a new migration.
  - `make migrate_up step=<number>`: Apply migrations with a specified step.
  - `make migrate_down step=<number>`: Roll back migrations with a specified step.
  - `make migrate_refresh`: Roll back all migrations and re-apply them.
  - `make migrate_force version=<version>`: Force a specific migration version.
  - `make run-cli`: Run the CLI application.

---

## Commands

<details>
<summary>### `make migration`</summary>

Creates a new migration file.

#### Description:
- Checks if the `name` variable is defined.
- If not, it displays an error message and exits.
- If defined, it creates a migration file in the specified directory with the given name.

</details>

<details>
<summary>### `make migrate_up`</summary>

Applies migrations.

#### Description:
- If `step` is not defined, it applies all pending migrations.
- If `step` is defined, it applies migrations up to the specified step.

</details>

<details>
<summary>### `make migrate_down`</summary>

Rolls back the last migration.

#### Description:
- If `step` is not defined, it rolls back the last migration.
- If `step` is defined, it rolls back the specified number of migrations.

</details>

<details>
<summary>### `make migrate_refresh`</summary>

Rolls back all migrations and re-applies them.

#### Description:
- Executes a command to roll back all migrations.
- Re-applies all migrations afterward.

</details>

<details>
<summary>### `make migrate_force`</summary>

Forces a specific migration version.

#### Description:
- Checks if the `version` variable is defined.
- If not, it displays an error message and exits.
- If defined, it forces the migration to the specified version.

</details>

<details>
<summary>### `make run-cli`</summary>

Runs the CLI application.

#### Description:
- Executes the command to run the CLI application.
- If the command fails, it does not halt the execution of the Makefile.

</details>

---

## Notes
- All migration commands utilize a PostgreSQL database connection string built from environment variables: `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, and `DB_NAME`.
- The `-verbose` flag is used for detailed output during migration operations.
