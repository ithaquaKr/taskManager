# Makefile
.PHONY: postgresql createdb dropdb force version migrateup migratedown run-linter tests runserver

# Environment variables for project
# ENV := $(PWD)/.env
# include $(ENV)

# Export all variable to sub-make
# export

# Internal variables
DB_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

#------------------------
# Migrate
#------------------------
force:
	migrate -path migrations --database "${DB_URL}" -verbose force $(VERSION)

version:
	migrate -path migrations --database "${DB_URL}" -verbose version $(VERSION)

migrateup:
	migrate -path migrations -database "${DB_URL}" -verbose up $(VERSION)

migratedown:
	migrate -path migrations -database "${DB_URL}" -verbose down $(VERSION)

#------------------------
# Database
#------------------------
postgresql:
	@echo "Running postgresql container..."
	docker run --name taskmanager_db -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d ${POSTGRES_IMAGE}

createdb:
	@echo "Creating database.."
	docker exec -it taskmanager_db createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	@echo "Droping database..."
	docker exec -it taskmanager_db dropdb ${POSTGRES_DB}


# ==============================================================================
# Tools commands

run-linter:
	@echo "Applying linter"
	golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

tests:
	go test -v -parallel 2 -cover ./...

#------------------------
# Run server
#------------------------

runserver:
	@echo "Running server..."
	go run main.go

