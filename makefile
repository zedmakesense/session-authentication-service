include .env
export

.PHONY: db-init db-drop db-reset

db-init:
	sudo -u postgres env PGPASSWORD=$(POSTGRES_PASSWORD) \
		psql \
		-v db_name=$(DB_NAME) \
		-v db_user=$(DB_USER) \
		-v db_password=$(DB_PASSWORD) \
		-f internal/db/bootstrap.sql

db-drop:
	dropdb --if-exists $(DB_NAME)

db-reset: db-drop db-init
