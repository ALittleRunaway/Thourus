include development/.env
export

# Don't forget to change them!
MIGRATION_NAME := test
FORCE_VERSION := 20220510111633

# proto generation

.PHONY: proto
proto:
	protoc	--go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative ./infrastructure/grpc/proto/cryptocore.proto

# compiling

.PHONY: build_linux
build_linux:
	env GOOS=linux GOARCH=amd64 go build -o thourus-api

.PHONY: build_mac
build_mac:
	env GOOS=darwin GOARCH=arm64 go build -o thourus-api

.PHONY: build_windows
build_windows:
	env GOOS=windows GOARCH=amd64 go build -o thourus-api

# migrations

.PHONY: migrate_new
migrate_new:
	migrate create -dir db/migration -ext sql ${MIGRATION_NAME}

.PHONY: migrate_up
migrate_up:
	migrate -source file://./db/migration -database mysql://${THOURUS_API_DB_USERNAME}:${THOURUS_API_DB_PASSWORD}@/thourus up

.PHONY: migrate_down
migrate_down:
	migrate -source file://./db/migration -database mysql://${THOURUS_API_DB_USERNAME}:${THOURUS_API_DB_PASSWORD}@/thourus down

.PHONY: migrate_version
migrate_version:
	migrate -source file://./db/migration -database mysql://${THOURUS_API_DB_USERNAME}:${THOURUS_API_DB_PASSWORD}@/thourus version

.PHONY: migrate_force
migrate_force:
	migrate -source file://./db/migration -database mysql://${THOURUS_API_DB_USERNAME}:${THOURUS_API_DB_PASSWORD}@/thourus force ${FORCE_VERSION}

