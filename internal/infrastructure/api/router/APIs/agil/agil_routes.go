package agil

import (
	controller "github.com/CASCOLOCO/boreal.git/internal/app/controllers/apis/agil"
	"github.com/labstack/echo/v4"
)

func AgilRoutes(e *echo.Echo) {
	group := e.Group("/api/agil/products")

	group.POST("/find", controller.GetAllProductsManual)

}
