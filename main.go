package main

import (
	"github.com/Dryluigi/golang-todos/controller"
	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	controller.NewGetAllTodosController(e, db)
	controller.NewCreateTodoController(e, db)
	controller.NewDeleteTodoController(e, db)
	controller.NewUpdateTodoController(e, db)
	controller.NewCheckTodoController(e, db)

	e.Start(":8080")
}
