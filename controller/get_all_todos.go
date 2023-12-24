package controller

import (
	"database/sql"
	"net/http"

	"github.com/Dryluigi/golang-todos/models"
	"github.com/labstack/echo"
)

type TodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func NewGetAllTodosController(e *echo.Echo, db *sql.DB) {
	e.GET("/todos", func(ctx echo.Context) error {
		user := ctx.Get("USER").(models.AuthJwtClaims)

		allowed := false
		for _, scope := range user.UserScopes {
			if scope == "todos:read" {
				allowed = true
				break
			}
		}
		if !allowed {
			return ctx.String(http.StatusForbidden, "Forbidden")
		}

		rows, err := db.Query("SELECT id, title, description, done FROM todos WHERE user_id = ?", user.UserId)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []TodoResponse
		for rows.Next() {
			var id int
			var title string
			var description string
			var done int

			err = rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			var todo TodoResponse
			todo.Id = id
			todo.Title = title
			todo.Description = description
			if done == 1 {
				todo.Done = true
			}

			res = append(res, todo)
		}

		return ctx.JSON(http.StatusOK, res)
	})
}
