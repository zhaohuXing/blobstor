package dao

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var dao *sql.DB

func Connect() *sql.DB {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	if dao != nil {
		return dao
	}
	dao = open()
	return dao
}

func open() *sql.DB {
	root := "root"
	password := "123456"
	url := "tcp(127.0.0.1:3306)/blobstor?charset=utf8"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s", root, password, url))
	if err != nil {
		log.Fatal("mysql connect failed.")
	}
	return db
}
