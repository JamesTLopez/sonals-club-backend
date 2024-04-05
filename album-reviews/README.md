# sonals-club-backend

### Creating migrations

Ensure that you are in the backend directory before running these commands.
Ensure that you have migrate CLI installed on your respective system, please refer to https://github.com/golang-migrate/migrate?tab=readme-ov-file

You can also take advantage of the Makefile (currenlt works only with localhost and not on docker environments as of Apr 4th,2024)

```
migrate create -ext sql -dir db/migrations -seq create_users_table
migrate -database "postgresql://james:secret@localhost:5432/sonalsguild?sslmode=disable" -path db/migrations up
```
