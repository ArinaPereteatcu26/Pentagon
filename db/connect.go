package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
)

var connectionURL = url.URL{
	Scheme: "sqlserver",
	User:   url.UserPassword(os.Getenv("DB_LOGIN"), os.Getenv("DB_PASS")),
	Host:   fmt.Sprintf("%s:%d", os.Getenv("DB_HOST"), 1433),
}

var db *sql.DB

func ConnectDB() error {
	var err error
	db, err = sql.Open("sqlserver", connectionURL.String())
	return err
}
