**Install sqlc:** go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

**RUN SQLC Generate**: sqlc generate

**Install migrate:** $ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.0/migrate.windows-amd64.zip

**RUN MIGRATE**: migrate create -ext sql -dir db/migration -seq init_schema

migrate -path ./db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migrate -path ./db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down