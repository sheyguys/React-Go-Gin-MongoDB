package main

import (
	"log"

	"github.com/sheyguys/gogin/route"

	"github.com/gin-gonic/gin"

	"github.com/globalsign/mgo"
)

const (
	mongoDBEnPint = "mongodb://localhost:27017"
	portWebServie = "localhost:8081"
)

func main() {
	connectionDB, err := mgo.Dial(mongoDBEnPint)
	if err != nil {
		log.Panic("Can no connect Database", err.Error())
	}
	router := gin.Default()
	route.NewRouteMember(router, connectionDB)
	router.Run(portWebServie)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
