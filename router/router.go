package router

import (
	"go_todo_app/model"
	"os"

	"github.com/labstack/echo/v4/middleware"

	"net/http"
	_ "net/http"

	"github.com/labstack/echo/v4"
)

// 関数 GetTasksHandlerは引数がecho.Context型のc で、戻り値はerror型である
func GetTasksHandler(c echo.Context) error {
    
    // model(package)の関数GetTasksを実行し、戻り値をtasks,errと定義する。
	tasks, err := model.GetTasks()
    
    // errが空でない時は StatusBadRequest(*5) を返す
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
    
    // StasusOK と tasksを返す
	return c.JSON(http.StatusOK, tasks)
}

//ルーティングを設定する関数 引数はecho.echo型のc であり、戻り値はerror型である
func SetRouter(e *echo.Echo) error {

    // 諸々の設定(*1)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
    
    
    // APIを書く場所
    e.GET("/api/tasks", GetTasksHandler)
    
    // 8000番のポートを開く(*2)
	err := e.Start(":8000")
	return err
}