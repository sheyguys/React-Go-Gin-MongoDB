package api

import (
	"log"
	"net/http"

	"github.com/sheyguys/gogin/model"
	"github.com/sheyguys/gogin/repository"

	"github.com/gin-gonic/gin"
)

type MemberAPI struct {
	MemberRepository repository.MemberRepository
}

func (api MemberAPI) MemberListHandler(context *gin.Context) {
	var memberInfo model.MemberInfo
	member, err := api.MemberRepository.GetAllMember()
	if err != nil {
		log.Println("error memberListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	memberInfo.Member = member
	context.JSON(http.StatusOK, memberInfo)
}

func (api MemberAPI) SaveMemberHandeler(context *gin.Context) {
	var member model.Member
	err := context.ShouldBindJSON(&member)
	if err != nil {
		log.Println("error AddMemberHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.MemberRepository.AddMember(member)
	if err != nil {
		log.Println("error AddMemberHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "susess"})
}

func (api MemberAPI) EditMemberNameHandler(context *gin.Context) {
	var member model.Member
	memberID := context.Param("member_id")
	err := context.ShouldBindJSON(&member)
	if err != nil {
		log.Println("error EditMemberNametHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.MemberRepository.EditMemberName(memberID, member)
	if err != nil {
		log.Println("error  EditMemberNametHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}
