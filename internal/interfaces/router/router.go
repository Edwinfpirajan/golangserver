package router

import (
	controller "github.com/CASCOLOCO/boreal.git/internal/app/controller"
	auth "github.com/CASCOLOCO/boreal.git/internal/interfaces/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoRoutes(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		// AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	group := e.Group("/api")
	groupProducts := e.Group("/products")

	//test
	groupProducts.GET("/all", controller.GetCombinationsAndProductOptionValues)
	groupProducts.GET("/p/all", controller.GetProduct)
	// groupProducts.GET("/combination/all", controller.CombinationsHandler)
	groupProducts.GET("/option/all", controller.GetProductOptionValues)
	// groupProducts.GET("/products/all", controller.GetCombinationsAndProductOptionValues)
	// groupProducts.GET("/all", controller.GetProducts)
	// LOGIN

	e.POST("/login", controller.Login)

	e.POST("/logout", controller.Logout, auth.OnlyAdmin)

	// USER ROUTES

	//ATTENDANCE REGISTER

	group.POST("/attendance/register", controller.SaveRegisterAttendance)
	group.GET("/attendance", controller.GetAllAttendance)
	group.GET("/attendance/validate/:pin", controller.ValidateEmploye)

	// ADMIN ROUTES

	//EMPLOYE MANAGE

	group.GET("/all", controller.GetAll)
	group.POST("/save", controller.SaveEmploye, auth.OnlyAdmin)
	group.DELETE("/delete/:id", controller.DeleteEmploye, auth.OnlyAdmin)
	group.GET("/find/:id", controller.GetEmploye, auth.OnlyAdmin)

	// SCHEDULE MANAGE

	group.GET("/schedule/all", controller.GetAllSchedule, auth.OnlyAdmin)
	group.POST("/schedule/save", controller.SaveSchedule, auth.OnlyAdmin)
	group.DELETE("/schedule/delete/:id", controller.DeleteSchedule, auth.OnlyAdmin)
}
