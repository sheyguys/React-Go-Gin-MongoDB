package main

import (
	"log"
	"github.com/gin-contrib/cors"
	"github.com/sheyguys/gogin/repository"

	"github.com/sheyguys/gogin/api"

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
	NewRouteMember(router, connectionDB)
	router.Run(portWebServie)
	
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, UPDATE")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }


func NewRouteMember(route *gin.Engine, connectionDB *mgo.Session) {
	memberRepository := repository.MemberRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	memberAPI := api.MemberAPI{
		MemberRepository: &memberRepository,
	}
	route.Use(cors.Default())
	route.GET("employee", memberAPI.MemberListHandler)
	route.POST("employee", memberAPI.AddMemberHandeler)
	route.PUT("employee/:member_id", memberAPI.EditMemberNameHandler)
	route.DELETE("employee/:member_id", memberAPI.DeleteMemberHandler)
}

