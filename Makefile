# Makefile

# Environment variables for project
ENV := $(PWD)/.env
include $(ENV)

# Export all variable to sub-make
export 

# Internal variables
DB_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable


#------------------------
# Database
#------------------------
postgresql:
	@echo "Running postgresql container..."
	docker run --name postgresql -p ${POSTGRES_PORT}:${POSTGRES_PORT} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d ${POSTGRES_IMAGE}

createdb:
	@echo "Creating database.."
	docker exec -it postgresql createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	@echo "Droping database..."
	docker exec -it postgresql dropdb ${POSTGRES_DB}

migrateup:
	@echo "Migrate up schema for database..."
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migratedown:
	@echo "Migrate down schema for database..."
	migrate -path db/migrations -database "${DB_URL}" -verbose down

.PHONY: postgresql createdb dropdb migrateup migratedown
