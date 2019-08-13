package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sheyguys/gogin/model"
)

type MemberRepository interface {
	GetAllMember() ([]model.Member, error)
	AddMember(member model.Member) error
	EditMemberName(memberID string, member model.Member) error
}

type MemberRepositoryMongo struct {
	ConnectionDB *mgo.Session
}

const (
	DBName     = "Employee"
	collection = "member"
)

func (memberMongo MemberRepositoryMongo) GetAllMember() ([]model.Member, error) {
	var member []model.Member
	err := memberMongo.ConnectionDB.DB(DBName).C(collection).Find(nil).All(&member)
	return member, err
}

func (memberMongo MemberRepositoryMongo) AddMember(member model.Member) error {
	return memberMongo.ConnectionDB.DB(DBName).C(collection).Insert(member)
}

func (memberMongo MemberRepositoryMongo) EditMemberName(memberID string, member model.Member) error {
	objectID := bson.ObjectIdHex(memberID)
	newName := bson.M{"$set": bson.M{"member_name": member.MemberName}}
	return memberMongo.ConnectionDB.DB(DBName).C(collection).UpdateId(objectID, newName)
}
