package router

import (
	"github.com/labstack/echo/v4"

	controller "github.com/CASCOLOCO/boreal.git/internal/app/controllers/employee"
	auth "github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/middleware"
)

func EmployeeManageRoutes(e *echo.Echo) {
	group := e.Group("/api/attendance")

	group.GET("/all", controller.GetAll)
	group.POST("/save", controller.SaveEmploye, auth.OnlyAdmin)
	group.DELETE("/delete/:id", controller.DeleteEmploye, auth.OnlyAdmin)
	group.GET("/find/:id", controller.GetEmploye, auth.OnlyAdmin)
}
