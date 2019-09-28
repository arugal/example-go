package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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
	route := gin.Default()

	dao := route.Group("/dao")

	{
		dao.GET("/select", Select)
		dao.POST("/insert", Insert)
		dao.POST("/update", Update)
		dao.POST("/delete", Delete)
	}

	route.Run(":9000")
}

func Select(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, users)
}

func Insert(ctx *gin.Context) {
	var user User
	if ctx.ShouldBind(&user) == nil {
		Exec(ctx, "INSERT INTO user (name, age) VALUES (?, ?)", func(result sql.Result) {
			id, err := result.LastInsertId()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"action": "LastInsertId",
					"name":   user.Name,
					"age":    user.Age,
				})
				log.Println(err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"id": id,
			})
		}, func() []interface{} {
			return []interface{}{user.Name, user.Age}
		})
	}
}

func Update(ctx *gin.Context) {
	var user User
	if ctx.ShouldBind(&user) == nil {
		Exec(ctx, "UPDATE user SET name = ?, age = ? WHERE  id = ?", func(result sql.Result) {
			num, err := result.RowsAffected()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"action": "RowsAffected",
					"id":     user.Id,
					"name":   user.Name,
					"age":    user.Age,
				})
				log.Println(err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"num": num,
			})
		}, func() []interface{} {
			return []interface{}{user.Name, user.Age, user.Id}
		})
	}
}

func Delete(ctx *gin.Context) {
	var user User
	if ctx.ShouldBind(&user) == nil {
		Exec(ctx, "DELETE FROM user where id = ?", func(result sql.Result) {
			num, err := result.RowsAffected()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"action": "RowsAffected",
					"id":     user.Id,
				})
				log.Println(err)
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"num": num,
			})
		}, func() []interface{} {
			return []interface{}{user.Id}
		})
	}
}

func Exec(ctx *gin.Context, sql string, resFunc func(result sql.Result), argsFunc func() []interface{}) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"action": "Prepare",
			"sql":    sql,
		})
		log.Printf("Prepare sql:%s err:%v", sql, err)
		return
	}

	args := argsFunc()
	res, err := stmt.Exec(args...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"action": "Exec",
			"args":   args,
		})
		log.Printf("Exec sql:%s args:%v err:%v", sql, args, err)
		return
	}
	resFunc(res)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
		panic(err)
	}
}

type User struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}
