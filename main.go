package main

import (
	"fmt"
	"github.com/labstack/echo"
	"test/Handler"
)

func main() {
	fmt.Println("Hello Echo!")
	e := echo.New()

	e.GET("/", Handler.HomeHandler)
	e.GET("/user", Handler.UserListHandler)
	e.POST("/user", Handler.AddUser)

	e.Logger.Fatal(e.Start(":9000"))

}
