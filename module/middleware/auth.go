package middlewares

import (
	"strings"

	"github.com/agusheryanto182/go-inventory-management/module/feature/staff"
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
				c.Logger().Error("unauthorized: missing token or invalid token")
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: missing token or invalid token")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				c.Logger().Error("unauthorized: invalid token " + err.Error())
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: invalid token "+err.Error())
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				c.Logger().Error("unauthorized: token is expired")
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: token is expired")
			}

			staffID, ok := claims["id"].(string)
			if !ok {
				c.Logger().Error("unauthorized: id user not valid || id : ", staffID)
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: id user not valid")
			}

			staff, err := staffService.GetStaffByID(staffID)
			if err != nil {
				c.Logger().Error("unauthorized: id is not found " + err.Error())
				return response.SendStatusUnauthorizedResponse(c, "unauthorized: id is not found "+err.Error())
			}

			c.Set("CurrentStaff", staff)

			return next(c)
		}
	}
}
