package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//"errors"
)

type user struct{
    ID int `json:"id"`
    Name string `json:"name"`
}

var users = []user{
    {ID:1, Name:"aa"},
    {ID:2, Name:"ab"},
}

func getUsers(c *gin.Context){
    c.IndentedJSON(http.StatusOK, users)
}

func main(){
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    fmt.Printf("mysql connection: %s \n", os.Getenv("MYSQL"))
    router := gin.Default()
    router.GET("/users", getUsers)
    router.Run("localhost:8080")
}
