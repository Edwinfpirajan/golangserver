package cascoloco

import (
	controller "github.com/CASCOLOCO/boreal.git/internal/app/controllers/apis/cascoloco"
	"github.com/labstack/echo/v4"
)

func CascolocoRoutes(e *echo.Echo) {
	group := e.Group("/api/products")

	group.GET("/all", controller.GetCombinationsAndProductOptionValues)
	group.GET("/p/all", controller.GetProduct)
	group.GET("/option/all", controller.GetProductOptionValues)
}
