package auth

import (
	"context"
	"fmt"
	"go-struct/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository interface {
	GetMigrations() []byte
}

type repository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &repository{db}
}

func (r *repository) GetMigrations() []byte {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM migrations")
	if err != nil {
		fmt.Printf("Error making register query: %s \n", err)
	}

	res := utils.PgSqlRowsToJson(rows)
	return res
}
