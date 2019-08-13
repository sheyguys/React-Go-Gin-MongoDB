package model

import (
	"github.com/globalsign/mgo/bson"
)

type MemberInfo struct {
	Member []Member `json:"member"`
}

type Member struct {
	MemberID      bson.ObjectId `json:"mamber_id" bson:"_id,omitempty"`
	MemberNameeng string        `json:"member_name_eng" bson:"member_name_eng"`
	MemberNameth  string        `json:"member_name_th" bson:"member_name_th"`
	MemberIdcard  string        `json:"member_idcard" bson:"member_idcard"`
	MemberPhone   string        `json:"member_phone" bson:"member_phone"`
	MemberAddress string        `json:"member_address" bson:"member_address"`
	MemberEmail   string        `json:"member_email" bson:"member_email"`
}
