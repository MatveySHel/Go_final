package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewPostgres(username, password, host, port, dbName string) (*pgx.Conn, error) {
    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
    conn, err := pgx.Connect(context.Background(), connStr)
    if err!=nil{
        return nil,err
    }
    return conn, nil
}