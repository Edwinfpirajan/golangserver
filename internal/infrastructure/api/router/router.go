package router

import (
	"github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/router/APIs/agil"
	"github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/router/APIs/cascoloco"
	"github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/router/admin"
	employee "github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/router/employee"

	"github.com/labstack/echo/v4"
)

func GlobalRouter(e *echo.Echo) {
	//Login Admin
	admin.GlobalRoutes(e)

	//Employee Register Attendance
	employee.AttendanceRoutes(e)

	//Employee Manage
	employee.EmployeeManageRoutes(e)
	employee.ScheduleManageRoutes(e)

	//Cascoloco Api
	cascoloco.CascolocoRoutes(e)

	//Agil api
	agil.AgilRoutes(e)
}
