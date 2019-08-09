package repository

import (
	"github.com/sheyguys/gogin/model"

	"github.com/globalsign/mgo"
)

type HeroRepository interface {
	GetAllHero() ([]model.Hero, error)
}

type HeroRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

const (
	DBName     = "DOTA2"
	collection = "Hero"
)

func (heroMongo HeroRepositoryMongo) GetAllHero() ([]model.Hero, error) {
	var hero []model.Hero
	err := heroMongo.ConnectionDB.DB(DBName).C(collection).Find(nil).All(&hero)
	return hero, err
}