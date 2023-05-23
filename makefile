db_dsn :=$(shell bin/dsngetter)

.PHONY: migrate_up migrate_down migrate_up_base migrate_down_base build_dsn_getter

migrate_up: build_dsn_getter migrate_up_base

migrate_down: build_dsn_getter migrate_down_base

migrate_up_base:
	@goose -dir migration -table goose_db_version postgres "$(db_dsn)" up

migrate_down_base:
	@goose -dir migration -table goose_db_version postgres "$(db_dsn)" down

build_dsn_getter:
	@go build -o bin/dsngetter cmd/dsn_getter/main.go