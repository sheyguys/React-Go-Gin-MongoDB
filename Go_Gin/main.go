package main

import (
	"github.com/sheyguys/gogin/route"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/globalsign/mgo"
)

const (
	mongoDBEnPint  = "mongodb://localhost:27017"
	portWebServie = ":3000"
)

func main() {
	connectionDB, err := mgo.Dial(mongoDBEnPint)
	if err != nil {
		log.Panic("Can no connect Database", err.Error())
	}
	router := gin.Default()
	route.NewRouteHero(router, connectionDB)
	router.Run(portWebServie)
}