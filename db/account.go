package db

import (
	"database/sql"
	"log"
	"main/model"
	"time"
)

type AccountDB struct {
	DB *sql.DB
}

func (receiver AccountDB) Create(account model.Account) error {
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

	_, err = stmt.Exec(account.PK, account.SK, account.OpenDate.Format("2006-01-02 15:04:05"), account.Type)
	return err
}

func (receiver AccountDB) GetAccount() (model.Account, error) {
	stmt, err := receiver.DB.Prepare("SELECT * FROM account;")
	if err != nil {
		return model.Account{}, err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Printf("stmt.Close() error: %v", err)
		}
	}(stmt)

	var openDate string
	var account model.Account
	if err := stmt.QueryRow().Scan(&account.PK, &account.SK, &openDate, &account.Type); err != nil {
		return model.Account{}, err
	}

	account.OpenDate, err = time.Parse("2006-01-02 15:04:05", openDate)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}
