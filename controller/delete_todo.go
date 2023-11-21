package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func NewDeleteTodoController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		_, err := db.Exec(
			"DELETE FROM todos WHERE id = ?",
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
