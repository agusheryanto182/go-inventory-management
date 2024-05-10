package database

import (
	"fmt"

	"github.com/agusheryanto182/go-inventory-management/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func InitDatabase() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		config.GetString("DB_USERNAME"),
		config.GetString("DB_PASSWORD"),
		config.GetString("DB_HOST"),
		config.GetString("DB_PORT"),
		config.GetString("DB_NAME"),
		config.GetString("DB_PARAMS"),
	)

	db, err := sqlx.Connect("pgx", dsn)

	logrus.Info("connected to database")

	return db, err
}
