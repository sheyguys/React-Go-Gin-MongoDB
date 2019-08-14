package route

import (
	"github.com/sheyguys/gogin/api"
	"github.com/sheyguys/gogin/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)

func NewRouteMember(route *gin.Engine, connectionDB *mgo.Session) {
	memberRepository := repository.MemberRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	memberAPI := api.MemberAPI{
		MemberRepository: &memberRepository,
	}
	route.GET("employee", memberAPI.MemberListHandler)
	route.POST("employee", memberAPI.AddMemberHandeler)
	route.PUT("employee/:member_id", memberAPI.EditMemberNameHandler)
	route.DELETE("employee/:member_id", memberAPI.DeleteMemberHandler)
}
