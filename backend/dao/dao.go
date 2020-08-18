package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	DB *sql.DB
}

func NewDao() *Dao {
	return &Dao{
		DB: NewDB(),
	}
}

func NewDB() *sql.DB {
	mysql, err := sql.Open("mysql", "root:0508@tcp(localhost:3306)/video?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	mysql.SetMaxOpenConns(10)
	return mysql
}
