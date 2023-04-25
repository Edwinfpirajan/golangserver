package admin

import (
	controller "github.com/CASCOLOCO/boreal.git/internal/app/controllers/admin"
	auth "github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/middleware"
	"github.com/labstack/echo/v4"
)

func GlobalRoutes(e *echo.Echo) {
	e.POST("/login", controller.Login)
	e.POST("/logout", controller.Logout, auth.OnlyAdmin)
}
