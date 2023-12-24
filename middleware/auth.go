package middleware

import (
	"net/http"
	"strings"

	"github.com/Dryluigi/golang-todos/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Path() == "/auth/register" || ctx.Path() == "/auth/login" || ctx.Path() == "/users/:userId/scopes/:scopeId/assign" {
			return next(ctx)
		}

		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return ctx.String(http.StatusUnauthorized, "empty token")
		}

		authHeaderArr := strings.Split(authHeader, " ")
		if len(authHeaderArr) != 2 {
			return ctx.String(http.StatusUnauthorized, "invalid token")
		}
		if authHeaderArr[0] != "Bearer" {
			return ctx.String(http.StatusUnauthorized, "invalid bearer token")
		}

		tokenStr := authHeaderArr[1]

		var jwtClaims models.AuthJwtClaims
		token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims, func(t *jwt.Token) (interface{}, error) {
			return []byte("TEST"), nil
		})
		if err != nil {
			return ctx.String(http.StatusUnauthorized, err.Error())
		}
		if !token.Valid {
			return ctx.String(http.StatusUnauthorized, "token invalid")
		}

		ctx.Set("USER", jwtClaims)

		return next(ctx)
	}
}
