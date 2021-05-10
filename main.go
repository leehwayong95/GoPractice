package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"test/DB"
	"test/route"
)

func main() {
	DB.InitDB()

	//router 등록
	e := echo.New()
	route.Router(e)

	//Setting middleware
	e.Logger.Fatal(e.Start(":9000"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut,http.MethodPost,http.MethodDelete},
	}))
}
