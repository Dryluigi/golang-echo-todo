package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Dryluigi/golang-todos/models"
	"github.com/labstack/echo"
)

type CheckRequest struct {
	Done bool `json:"done"`
}

func NewCheckTodoController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		user := ctx.Get("USER").(models.AuthJwtClaims)
		id := ctx.Param("id")

		allowed := false
		for _, scope := range user.UserScopes {
			if scope == "todos:update" {
				allowed = true
				break
			}
		}
		if !allowed {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		var request CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var doneInt int
		if request.Done {
			doneInt = 1
		}

		_, err := db.Exec(
			"UPDATE todos SET done = ? WHERE id = ? AND user_id = ?",
			doneInt,
			id,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
