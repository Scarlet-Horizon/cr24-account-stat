package db

import (
	"database/sql"
	"log"
	"main/request"
	"main/response"
	"time"
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

	stmt, err = receiver.DB.Prepare("INSERT INTO stat (visited, fk_endpoint) VALUES (?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), id)
	return err
}

func (receiver StatDB) LastEndpoint() ([]response.EndpointLast, error) {
	stmt, err := receiver.DB.Prepare("SELECT e.name, s.visited FROM endpoint e" +
		" JOIN stat s on e.id_endpoint = s.fk_endpoint WHERE s.visited = (SELECT MAX(visited) FROM stat);")
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Printf("stmt.Close() error: %v", err)
		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close() error: %v", err)
		}
	}(rows)

	var last []response.EndpointLast

	for rows.Next() {
		var result response.EndpointLast

		if err := rows.Scan(&result.Name, &result.Visited); err != nil {
			log.Printf("rows.Scan() error: %v", err)
			continue
		}

		last = append(last, result)
	}
	if err := rows.Err(); err != nil {
		log.Printf("rows.Err() error: %v", err)
	}
	return last, nil
}

func (receiver StatDB) mostAndAll(query string) ([]response.EndpointStat, error) {
	stmt, err := receiver.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			log.Printf("stmt.Close() error: %v", err)
		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close() error: %v", err)
		}
	}(rows)

	var data []response.EndpointStat

	for rows.Next() {
		var result response.EndpointStat

		if err := rows.Scan(&result.Name, &result.Count); err != nil {
			log.Printf("rows.Scan() error: %v", err)
			continue
		}

		data = append(data, result)
	}
	if err := rows.Err(); err != nil {
		log.Printf("rows.Err() error: %v", err)
	}
	return data, nil
}

func (receiver StatDB) MostCalled() ([]response.EndpointStat, error) {
	return receiver.mostAndAll("SELECT e.name, COUNT(e.id_endpoint) count FROM endpoint e" +
		" JOIN stat s on e.id_endpoint = s.fk_endpoint" +
		" GROUP BY e.id_endpoint" + "" +
		" HAVING COUNT(e.id_endpoint) = (SELECT MAX(count) FROM endpointCount);")
}

func (receiver StatDB) EndpointAll() ([]response.EndpointStat, error) {
	return receiver.mostAndAll("SELECT * FROM endpointCount ORDER BY count DESC;")
}
