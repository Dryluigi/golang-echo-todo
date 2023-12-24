package main

import (
	"github.com/Dryluigi/golang-todos/controller"
	"github.com/Dryluigi/golang-todos/database"
	"github.com/Dryluigi/golang-todos/middleware"
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

	e.Use(middleware.AuthMiddleware)

	controller.NewGetAllTodosController(e, db)
	controller.NewCreateTodoController(e, db)
	controller.NewDeleteTodoController(e, db)
	controller.NewUpdateTodoController(e, db)
	controller.NewCheckTodoController(e, db)
	controller.NewRegisterController(e, db)
	controller.NewLoginController(e, db)
	controller.NewCreateScopeController(e, db)
	controller.NewDeleteScopeController(e, db)
	controller.NewAssignUserScopeController(e, db)

	e.Start(":8080")
}
