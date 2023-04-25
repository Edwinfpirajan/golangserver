package router

import (
	"github.com/labstack/echo/v4"

	controller "github.com/CASCOLOCO/boreal.git/internal/app/controllers/employee"
	auth "github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/middleware"
)

func ScheduleManageRoutes(e *echo.Echo) {
	group := e.Group("/api/schedule")

	group.GET("/all", controller.GetAllSchedule, auth.OnlyAdmin)
	group.POST("/save", controller.SaveSchedule, auth.OnlyAdmin)
	group.DELETE("/delete/:id", controller.DeleteSchedule, auth.OnlyAdmin)
}
