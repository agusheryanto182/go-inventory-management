package middlewares

import (
	"strings"

	utils "github.com/agusheryanto182/go-inventory-management/utils/jwt"
	"github.com/agusheryanto182/go-inventory-management/utils/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(jwtService utils.JWTInterface, staffService staff.ServiceStaffInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") {
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: invalid token "+err.Error())
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: token is expired "+err.Error())
			}

			staffIDFloat, ok := claims["id"].(float64)
			if !ok {
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: id user not valid "+err.Error())
			}

			staffID := uint64(staffIDFloat)

			staff, err := staffService.GetStaffById(staffID)
			if err != nil {
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: id is not found "+err.Error())
			}

			c.Set("CurrentStaff", staff)

			return next(c)
		}
	}
}
