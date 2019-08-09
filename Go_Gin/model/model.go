package model

import (
	"github.com/globalsign/mgo/bson"
)

type HeroInfo struct {
	Hero []Hero `json:"hero"`
}

type Hero struct {
	HeroID   bson.ObjectId `json:"hero_id" bson:"_id,omitempty"`
	HeroName string        `json:"hero_name" bson:"hero_name"`
	HeroType string        `json:"hero_type" bson:"hero_type"`
	Winrate  int           `json:"winrate" bson:"winrate"`
}
