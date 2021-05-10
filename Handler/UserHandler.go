package Handler

import (
	"github.com/labstack/echo"
	"net/http"
	"test/DB"
	"test/Model"
)


//UserListHandler 전체조회
func UserListHandler(c echo.Context) error {
	db := DB.DBManager()
	var users []Model.Push

	var response Model.DefaultResponse
	err := db.Order("createat desc").Find(&users).Error
	if err != nil {
		response.Status = false
		response.Reason = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Status = true
	response.Data = users

	return c.JSON(http.StatusOK, response)
}

//UserHandler is select user 생성일자조회
func UserHandler(c echo.Context) error {
	db := DB.DBManager()
	params := make(map[string]string)
	c.Bind(&params)
	var user Model.Push
	db.First(&user, params["pid"])
	db.Where(&user.ID).Find(&user.Createat)
	return c.JSON(http.StatusOK, user.Createat)
}

//AddUser return addUser Result
func AddUser(c echo.Context) error {
	db := DB.DBManager()
	params := make(map[string]string)
	c.Bind(&params)
	user := Model.Push{
		ID:   params["pid"],
		ProgressName: params["pName"],
		PushTitle:    params["pushTitle"],
		PushContents: params["pushContents"],
	}

	//DB 쿼리
	var response Model.DefaultResponse
	err := db.Create(&user).Error
	if err != nil {
		response.Status = false
		response.Reason = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}
	response.Status = true
	return c.JSON(http.StatusOK, response)
}
