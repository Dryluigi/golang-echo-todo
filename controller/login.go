package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Dryluigi/golang-todos/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func NewLoginController(e *echo.Echo, db *sql.DB) {
	e.POST("/auth/login", func(ctx echo.Context) error {
		var request LoginRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		row := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", request.Email)
		if row.Err() != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var retrievedId int
		var retrievedName, retrievedEmail, retrievedPassword string

		err := row.Scan(&retrievedId, &retrievedName, &retrievedEmail, &retrievedPassword)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ctx.String(http.StatusUnauthorized, "email not found")
			}

			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		err = bcrypt.CompareHashAndPassword([]byte(retrievedPassword), []byte(request.Password))
		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}

		rows, err := db.Query("SELECT scopes.name FROM users LEFT JOIN user_scopes ON user_scopes.user_id = users.id JOIN scopes ON user_scopes.scope_id = scopes.id WHERE email = ?", request.Email)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, row.Err().Error())
		}

		var scopes []string = make([]string, 0)
		for rows.Next() {
			var scopeName string
			err = rows.Scan(&scopeName)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			scopes = append(scopes, scopeName)
		}
		rows.Close()

		jwtClaim := models.AuthJwtClaims{
			UserId:     retrievedId,
			UserName:   retrievedName,
			UserEmail:  retrievedEmail,
			UserScopes: scopes,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
		tokenStr, err := token.SignedString([]byte("TEST"))
		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}

		response := LoginResponse{
			AccessToken: tokenStr,
		}

		return ctx.JSON(http.StatusOK, response)
	})

}
