package main

import (
	"go_todo_app/model"
	"go_todo_app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
	e := echo.New()
	router.SetRouter(e)
}