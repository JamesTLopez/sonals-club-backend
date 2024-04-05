include .env

create_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)
generate_migration:
	migrate -database "${DATABSE_URL_LOCAL}" -path db/migrations up
drop_migration:
	migrate -database "${DATABSE_URL_LOCAL}" -path db/migrations drop
