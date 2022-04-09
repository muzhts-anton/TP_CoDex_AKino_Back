package database

import (
	"codex/internal/pkg/utils/log"
	
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const connString = "user=akino password=1234 host=localhost port=5432 dbname=codex"

type DBrow_t [][]byte

type ConnectionPool interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

type DBManager struct {
	Pool ConnectionPool
}

func InitDatabase() *DBManager{
	return &DBManager{Pool: nil}
}

func (dbm *DBManager)Connect()  {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Warn("Postgres error")
		log.Error(err)
		return
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Warn("Ping error")
		log.Error(err)
		return
	}
	
	log.Info("Successful connection to postgres")
	dbm.Pool = pool
}

func (dbm *DBManager) Disconnect() {
	dbm.Pool.Close()
	log.Info("Postgres disconnected")
}

func (dbm *DBManager) Query(queryString string, params ...interface{}) ([]DBrow_t, error) {
	transactionContext := context.Background()
	tx, err := dbm.Pool.Begin(transactionContext)
	if err != nil {
		log.Warn("Error connecting to a pool")
		log.Error(err)
		return nil, err
	}
	
	defer tx.Rollback(transactionContext)

	rows, err := tx.Query(transactionContext, queryString, params...)
	if err != nil {
		log.Warn("Error in query: " + queryString)
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]DBrow_t, 0)
	for rows.Next() {
		rowBuffer := make(DBrow_t, 0)
		rowBuffer = append(rowBuffer, rows.RawValues()...)
		result = append(result, rowBuffer)
	}

	err = tx.Commit(transactionContext)
	if err != nil {
		log.Warn("Error committing")
		log.Error(err)
		return nil, err
	}

	return result, nil
}

func (dbm *DBManager) Execute(queryString string, params ...interface{}) error {
	transactionContext := context.Background()
	tx, err := dbm.Pool.Begin(transactionContext)
	if err != nil {
		log.Warn("Error connecting to a pool")
		log.Error(err)
		return err
	}

	defer tx.Rollback(transactionContext)

	_, err = tx.Exec(transactionContext, queryString, params...)
	if err != nil {
		log.Warn("Error in query: " + queryString)
		log.Error(err)
		return err
	}

	err = tx.Commit(transactionContext)
	if err != nil {
		log.Warn("Error committing")
		log.Error(err)
		return err
	}

	return nil
}
