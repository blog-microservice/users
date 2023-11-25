package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/saufiroja/blog-microservice/users/config"
	"github.com/saufiroja/blog-microservice/users/utility"
)

type IApplicationDatabase interface {
	StartTransaction() (*sql.Tx, error)
	CommitTransaction(tx *sql.Tx) error
	RollbackTransaction(tx *sql.Tx) error
}

type ApplicationDatabase struct {
	db *sql.DB
}

func NewApplicationDatabase(conf *config.AppConfig, logger utility.CustomLogger) IApplicationDatabase {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	name := conf.Postgres.Name
	ssl := conf.Postgres.SSL

	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, name, ssl)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		logger.Error(fmt.Sprintf("error connecting to database: %s", err.Error()))
		panic(err)
	}

	// ping database
	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("error pinging database: %s", err.Error()))
		panic(err)
	}

	// connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60)
	db.SetConnMaxIdleTime(10)

	logger.Info("database connected")

	return &ApplicationDatabase{db}
}

func (a *ApplicationDatabase) StartTransaction() (*sql.Tx, error) {
	return a.db.Begin()
}

func (a *ApplicationDatabase) CommitTransaction(tx *sql.Tx) error {
	return tx.Commit()
}

func (a *ApplicationDatabase) RollbackTransaction(tx *sql.Tx) error {
	return tx.Rollback()
}
