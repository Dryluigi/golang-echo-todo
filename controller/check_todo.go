package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type CheckRequest struct {
	Done bool `json:"done"`
}

func NewCheckTodoController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var doneInt int
		if request.Done {
			doneInt = 1
		}

		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ?",
			doneInt,
			id,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
