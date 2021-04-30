package Handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	PushModel "test/Model"
)

var UserList []PushModel.Push

func UserListHandler(c echo.Context) error {
	if len(UserList) > 0 {
		fmt.Println(UserList)
		return c.JSON(http.StatusOK, UserList)
	}
	return c.String(http.StatusOK, "No member")
}

func AddUser(c echo.Context) error {
	params := make(map[string]string)
	c.Bind(&params)
	user := PushModel.Push{
		ProgressId:   params["pid"],
		ProgressName: params["pName"],
		PushTitle:    params["pushTitle"],
		PushContents: params["pushContents"],
	}
	UserList = append(UserList, user)
	return c.JSON(http.StatusOK, "result: success")
}
