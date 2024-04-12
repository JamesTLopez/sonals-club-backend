include .env

create_migration:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)
generate_migration:
	migrate -database "${DATABSE_URL_LOCAL}" -path internal/db/migrations up
drop:
	migrate -database "${DATABSE_URL_LOCAL}" -path internal/db/migrations drop
