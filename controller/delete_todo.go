package controller

import (
	"database/sql"
	"net/http"

	"github.com/Dryluigi/golang-todos/models"
	"github.com/labstack/echo"
)

func NewDeleteTodoController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		user := ctx.Get("USER").(models.AuthJwtClaims)

		allowed := false
		for _, scope := range user.UserScopes {
			if scope == "todos:delete" {
				allowed = true
				break
			}
		}
		if !allowed {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		id := ctx.Param("id")

		_, err := db.Exec(
			"DELETE FROM todos WHERE id = ? AND user_id = ?",
			id,
			user.UserId,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})
}
