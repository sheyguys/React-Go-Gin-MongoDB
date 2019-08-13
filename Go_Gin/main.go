package main

import (
	"log"

	"github.com/sheyguys/gogin/route"

	"github.com/gin-gonic/gin"

	"github.com/globalsign/mgo"
)

const (
	mongoDBEnPint = "mongodb://localhost:27017"
	portWebServie = ":4000"
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
