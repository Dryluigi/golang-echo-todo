package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type CreateScopeRequest struct {
	Name string `json:"name"`
}

func NewCreateScopeController(e *echo.Echo, db *sql.DB) {
	e.POST("/scopes", func(ctx echo.Context) error {
		var request CreateScopeRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		row := db.QueryRow("SELECT id FROM scopes WHERE name = ?", request.Name)
		if row.Err() != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var retrievedId int
		err := row.Scan(&retrievedId)
		if err == nil {
			return ctx.String(http.StatusUnauthorized, "duplicate scope found")
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		_, err = db.Exec(
			"INSERT INTO scopes (name) VALUES (?)",
			request.Name,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

}
