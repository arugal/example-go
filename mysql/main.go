package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// https://blog.51cto.com/zhixinhu/1844734
// http://www.01happy.com/golang-mysql-demo/

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(192.168.2.124:3306)/test?charset=utf8")
	checkErr(err, "db init err")
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	db.Stats()
}

func main() {
	http.HandleFunc("/select", Select)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

func Select(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select id, `name`, age from user")
	checkErr(err, "db query err")
	defer rows.Close()

	_, _ = rows.Columns()
	var users []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		checkErr(err, "rows scan err")
		users = append(users, user)
	}

	result, err := json.Marshal(users)
	checkErr(err, "json Marshal users err")
	_, _ = w.Write(result)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
		panic(err)
	}
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
