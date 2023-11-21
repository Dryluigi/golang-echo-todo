package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateTodoController(e *echo.Echo, db *sql.DB) {
	e.POST("/todos", func(ctx echo.Context) error {
		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

}
