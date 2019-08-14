package api_test

import (
	"bytes"
	"github.com/sheyguys/gogin/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func Test_MemberListHandler_Should_Be_MemberInfo(t *testing.T) {
	expected := listMember
	request := httptest.NewRequest("GET", "/emplopyee", nil)
	writer := httptest.NewRecorder()
	memberAPI := api.MemberAPI{
		MemberRepository: &mockMemberRepository{},
	}

	testRoute := gin.Default()
	testRoute.GET("employee", memberAPI.MemberListHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualRespone, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actualRespone))
}

func Test_AddMemberHandler_Input_Member_Shoud_Be_Created(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	member := []byte(`{"member_name_eng":"Topson Noland","member_name_th":"ท็อปสัน โนแลนด์","member_idcard":1449900795130,"member_phone":"0884167209",
	"member_address":"Suranaree University","member_email":"b6020911@gmail.com","member_facebook":"Topson Allstars"}`)
	request := httptest.NewRequest("POST", "/employee", bytes.NewBuffer(member))
	writer := httptest.NewRecorder()
	memberAPI := api.MemberAPI{
		MemberRepository: &mockMemberRepository{},
	}

	testRoute := gin.Default()
	testRoute.POST("employee", memberAPI.AddMemberHandeler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualStatusCode := response.StatusCode

	assert.Equal(t, expectedStatusCode, actualStatusCode)
}

func Test_EditMemberNamethHandler_Input_Name_Chatcha_Shoukd_Be_Edited(t *testing.T) {
	expectedStatusCode := http.StatusOK
	member := []byte(`{"member_name_th":"ฌัชชา คำมูล"}`)
	request := httptest.NewRequest("PUT", "/employee/5beaf7bd62e63844ce22cc58", bytes.NewBuffer(member))
	writer := httptest.NewRecorder()
	memberAPI := api.MemberAPI{
		MemberRepository: &mockMemberRepository{},
	}

	testRoute := gin.Default()
	testRoute.PUT("employee/:member_id", memberAPI.EditMemberNameHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualStatusCode := response.StatusCode

	assert.Equal(t, expectedStatusCode, actualStatusCode)
}

func Test_DeleteMemberByIDHandler_Input_Id_5beaf7bd62e63844ce22cc57_Shout_Be_Delete_Member_Name_Kasinan(t *testing.T) {
	expectedStatusCode := http.StatusNoContent
	request := httptest.NewRequest("DELETE", "/employee/5beaf7bd62e63844ce22cc57", nil)
	writer := httptest.NewRecorder()
	memberAPI := api.MemberAPI{
		MemberRepository: &mockMemberRepository{},
	}

	testRoute := gin.Default()
	testRoute.DELETE("employee/:member_id", memberAPI.DeleteMemberHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualStatusCode := response.StatusCode

	assert.Equal(t, expectedStatusCode, actualStatusCode)
}