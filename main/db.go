package main

import (
	"database/sql"
	"log"
	"github.com/lib/pq"
)
func connect () *sql.DB{
	connector, err:=pq.NewConnector("host=localhost port=5432 user=postgres dbname=lesson_db sslmode=disable")
	if err:=nil{
		log.Println(err)
	}
	return sql.OpenDB(connector)
}