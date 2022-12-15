package db

import (
	"database/sql"
	"log"
	"main/request"
	"time"
)

type AccountDB struct {
	DB *sql.DB
}

func (receiver AccountDB) Create(account request.Account) error {
	stmt, err := receiver.DB.Prepare("DELETE FROM account;")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Printf("stmt.Close() error: %v", err)
		}
	}(stmt)

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	stmt, err = receiver.DB.Prepare("INSERT INTO account VALUES (?, ?, ?, ?);")

	_, err = stmt.Exec(account.PK, account.SK, time.Now().Format("2006-01-02 15:04:05"), account.Type)
	return err
}
