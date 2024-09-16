package main

import (
	"shared"

	"github.com/gin-gonic/gin"
)

func main() {
	db := shared.ConnectDB()

	r := gin.Default()

	r.Run(":8082")
}
