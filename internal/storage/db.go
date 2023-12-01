package storage

import (
	"DB_project/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func New(cfg *config.Config) error {
	var err error
	const op = "storage.postgres.New"

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	fmt.Println(psqlconn)

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func Ping() error {
	return db.Ping()
}

func Close() error {
	return db.Close()
}
