package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const connString = `
user=akino
password=1234
host=localhost
port=5432
dbname=codex
`

type DBbyterow [][]byte

type ConnectionPool interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

type DBManager struct {
	Pool ConnectionPool
}

func InitDatabase() *DBManager {
	return &DBManager{
		Pool: nil,
	}
}

func (dbm *DBManager) Connect() {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
		return
	}

	dbm.Pool = pool
}

func (dbm *DBManager) Disconnect() {
	dbm.Pool.Close()
}

func (dbm *DBManager) Query(queryString string, params ...interface{}) ([]DBbyterow, error) {
	transactionContext := context.Background()
	tx, err := dbm.Pool.Begin(transactionContext)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer tx.Rollback(transactionContext)

	rows, err := tx.Query(transactionContext, queryString, params...)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]DBbyterow, 0)
	for rows.Next() {
		rowBuffer := make(DBbyterow, 0)
		rowBuffer = append(rowBuffer, rows.RawValues()...)
		result = append(result, rowBuffer)
	}

	err = tx.Commit(transactionContext)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return result, nil
}
