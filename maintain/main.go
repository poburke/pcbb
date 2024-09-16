package main

import (
	"log"
	"shared"

	"github.com/gin-gonic/gin"
)

func main() {
	db := shared.ConnectDB()
	redisClient := shared.ConnectRedis()

	r := gin.Default()

	log.Println("Maintain service running...")
	r.Run(":8081")
}
