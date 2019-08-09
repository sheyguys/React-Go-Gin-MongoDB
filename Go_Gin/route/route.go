package route

import (
	"github.com/sheyguys/gogin/api"
	"github.com/sheyguys/gogin/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)

func NewRouteHero(route *gin.Engine, connectionDB *mgo.Session) {
	heroRepository := repository.HeroRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	heroAPI := api.HeroAPI{
		HeroRepository: &heroRepository,
	}
	route.GET("dota2", heroAPI.HeroListHandler)
}