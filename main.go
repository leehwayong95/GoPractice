package main

import (
	"fmt"
	"github.com/labstack/echo"
	"test/DB"
	"test/route"
)

func main() {
	fmt.Println("Initial DB....")
	DB.InitDB()

	//router 등록
	e := echo.New()
	route.Router(e)
	e.Logger.Fatal(e.Start(":9000"))
}
