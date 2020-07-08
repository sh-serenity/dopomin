 
package main

import (
	"fmt"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
func dbConnect() (db *sql.DB) {
    db, err := sql.Open("mysql", "dbuser:passblyadovo@tcp(dbhost:3306)/db")
    if err != nil {
	fmt.Println(err)
    }
return db
}