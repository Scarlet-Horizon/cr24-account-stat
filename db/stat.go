package db

import (
	"database/sql"
	"errors"
	"log"
	"main/request"
)

type StatDB struct {
	DB *sql.DB
}

func (receiver StatDB) CreateStat(statRequest request.StatRequest) error {
	stmt, err := receiver.DB.Prepare("SELECT id_endpoint FROM endpoint WHERE name = ?;")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Printf("stmt.Close() error: %v", err)
		}
	}(stmt)

	var id int
	if err := stmt.QueryRow(statRequest.Endpoint).Scan(&id); err != nil {
		return err
	}

	if id == 0 {
		return errors.New("invalid endpoint")
	}

	stmt, err = receiver.DB.Prepare("INSERT INTO stat (fk_endpoint) VALUES (?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
