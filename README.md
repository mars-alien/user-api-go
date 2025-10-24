# user-api-go

A small Go REST API that manages users (name and date-of-birth). Built with Fiber, PostgreSQL and sqlc for type-safe DB access.

## Features
- CRUD for users
- Fiber HTTP server
- SQL migrations in `db/migrations`
- sqlc configuration in `sqlc.yaml` to generate typed DB code

## Prerequisites
- Go 1.24 or later
- PostgreSQL (local or remote)
- sqlc (optional — only required if you edit SQL and need to regenerate code)
- (Optional) Docker & Docker Compose for running Postgres locally

## Quick start (recommended)

1) Clone repository (if not already):

```cmd
cd /d %HOMEDRIVE%\path\to\where\you\want\it
git clone https://github.com/mars-alien/user-api-go.git
cd user-api-go
```

2) Create a `.env` file (the project uses `github.com/joho/godotenv` so `.env` is loaded automatically). Example `.env`:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=user-api
SERVER_PORT=3000
```

3) Start a local Postgres (if you don't have one). Quick Docker example:

```cmd
docker run --name user-api-postgres -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -e POSTGRES_DB=user-api -p 5432:5432 -d postgres:15
```

4) Run the DB migration to create the `users` table:

```cmd
psql -h %DB_HOST% -p %DB_PORT% -U %DB_USER% -d %DB_NAME% -f db/migrations/001_create_users_table.sql
```

On Windows you may run the `psql` command from Git Bash, WSL or use a GUI client (pgAdmin, DBeaver) to execute `db/migrations/001_create_users_table.sql`.

5) (Optional) Generate sqlc code if you change SQL files. Install sqlc then run:

```cmd
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
sqlc generate
```

6) Build and run the server

```cmd
go build -o bin/user-api ./cmd/server
bin\user-api
```

Or run directly:

```cmd
go run ./cmd/server
```

Server listens on the port defined by `SERVER_PORT` (default `3000`). Example: `http://localhost:3000`.

## API Endpoints
Base path: `/users`

- Create user

  POST /users

  Request JSON body:
  ```json
  {
    "name": "Alice",
    "dob": "1990-01-20"
  }
  ```

  Example (cmd.exe + PowerShell compatible curl):
  ```cmd
  curl -X POST http://localhost:3000/users -H "Content-Type: application/json" -d "{\"name\":\"Alice\",\"dob\":\"1990-01-20\"}"
  ```

- Get user

  GET /users/:id

  Example:
  ```cmd
  curl http://localhost:3000/users/1
  ```

- List users

  GET /users?page=1&page_size=10

- Update user

  PUT /users/:id

  Body (partial allowed depending on model):
  ```json
  {
    "name": "Alice Smith",
    "dob": "1990-01-20"
  }
  ```

  Example:
  ```cmd
  curl -X PUT http://localhost:3000/users/1 -H "Content-Type: application/json" -d "{\"name\":\"Alice Smith\",\"dob\":\"1990-01-20\"}"
  ```

- Delete user

  DELETE /users/:id

  Example:
  ```cmd
  curl -X DELETE http://localhost:3000/users/1
  ```

## Environment variables
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` — used to form the Postgres connection string
- `SERVER_PORT` — port the HTTP server listens on (default: `3000`)

These are read from environment (or `.env`) in `config/config.go`.

## Development notes
- The project uses `sqlc` to generate typed DB code from SQL files in `db/queries`. If you modify SQL, run `sqlc generate` and commit the generated files under `db/sqlc`.
- Logging is handled by `go.uber.org/zap`; middleware logs requests.




## Troubleshooting
- `psql` connection errors: confirm Postgres is running, check `pg_hba.conf` and firewall rules, and ensure your `.env` values are correct.
- Line ending warnings when committing (LF -> CRLF): on Windows set `git config --global core.autocrlf true` to normalize.

## Contributing
PRs welcome. If you change SQL queries, remember to run `sqlc generate` and include generated output.

## License
This project doesn't include a license file. Add one (e.g., MIT) if you plan to publish.
