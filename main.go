package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nxtCoder36/graphql-golang-server/handler"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/todos", handler.TodoGraphRouter)
	port := os.Getenv("PORT")
	fmt.Println("Server Running on port:", port)
	router.Run(":" + port)
}
