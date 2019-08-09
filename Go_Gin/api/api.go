package api

import (
	"github.com/sheyguys/gogin/model"
	"github.com/sheyguys/gogin/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HeroAPI struct {
	HeroRepository repository.HeroRepository
}

func (api HeroAPI) HeroListHandler(context *gin.Context) {
	var heroInfo model.HeroInfo
	hero, err := api.HeroRepository.GetAllHero()
	if err != nil {
		log.Println("error heroListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	heroInfo.Hero = hero
	context.JSON(http.StatusOK, heroInfo)
}