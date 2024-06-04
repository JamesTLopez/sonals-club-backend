package services

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

var db *sql.DB
const dbTimeout = time.Second * 3
var psql sq.StatementBuilderType

type Models struct {
	Songs Song
	User User
	Samples Sample
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return Models{}
}