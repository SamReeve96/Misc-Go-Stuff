package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //Needed to use mysql
)

func main() {
	db := connectToDB()
	ReturnStr := queryDb(db, "firstname", "user", "id = 1")
	fmt.Println(ReturnStr)
}

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/documize")
	if err != nil {
		log.Println(err)
	}
	return db
}

func queryDb(db *sql.DB, fieldName string, tableName string, searchCriteria string) string {
	var ReturnStr string
	queryString := ("select " + fieldName + " from " + tableName + " where " + searchCriteria)
	err := db.QueryRow(queryString).Scan(&ReturnStr)
	//fmt.Println(queryString)
	//err := db.QueryRow("select firstname from user where id = 1").Scan(&ReturnStr)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	return ReturnStr
}
