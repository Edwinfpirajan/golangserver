package main

import (
	"github.com/CASCOLOCO/boreal.git/internal/infrastructure/api/router"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	router.GlobalRouter(e)

	e.Logger.Fatal(e.Start(":8080"))

	// PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	// HOST := os.Getenv("SERVER_HOST")

	// e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", HOST, PORT)))
}
