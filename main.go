package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "database/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    "github.com/go-sql-driver/mysql"
	//"errors"
)

type user struct{
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

var test_users = []user{
    {ID:1, Name:"aa", Email:"aa@gmail.com"},
    {ID:2, Name:"ab", Email:"ab@gmail.com"},
}

var db_users []user

func getUsers(c *gin.Context){
    if len(db_users) == 0 {
        fmt.Printf("\033[33mGetting users from test data!\033[0m\n")
        c.IndentedJSON(http.StatusOK, test_users)
    } else {
        fmt.Printf("\033[33mGetting users from mysql database!\033[0m\n")
        c.IndentedJSON(http.StatusOK, db_users)
    }
}

func main(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal("\033[31mError loading .env file\033[0m\n")
    }
    cfg := mysql.Config{
        User: os.Getenv("MYSQL_USER"),
        Passwd: os.Getenv("MYSQL_PASS"),
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "livebetter",
    }
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)    
    }
    fmt.Printf("\033[33mDB Opened!\033[0m\n")
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Printf("\033[36mDB Connected!\033[0m\n")
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        log.Fatal("\033[31mQuery Error!\033[31m\n")
    }
    defer rows.Close()
    for rows.Next() {
        var p user
        if err := rows.Scan(&p.ID,&p.Name,&p.Email); err != nil {
            log.Fatal("\033[31mScaning Row Error!\033[31m\n")
        }
        db_users = append(db_users, p)
    }
    router := gin.Default()
    router.GET("/users", getUsers)
    router.Run("localhost:8080")
}
