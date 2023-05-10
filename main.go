package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
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
    router := gin.Default()
    router.GET("/users", getUsers)
    router.Run("localhost:8080")
}
