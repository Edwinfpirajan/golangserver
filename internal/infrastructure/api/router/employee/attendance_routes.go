package router

import (
	"github.com/labstack/echo/v4"

	controller "github.com/CASCOLOCO/boreal.git/internal/app/controllers/employee"
)

func AttendanceRoutes(e *echo.Echo) {
	group := e.Group("/api/attendance")

	group.GET("/all", controller.GetAllAttendance)
	group.GET("/validate/:pin", controller.ValidateEmploye)
	group.POST("/register", controller.SaveRegisterAttendance)
}
