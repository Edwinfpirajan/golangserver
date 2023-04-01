package main

import (
	"net"
	"net/http"

	"github.com/CASCOLOCO/boreal.git/internal/interfaces/router"
	"github.com/labstack/echo/v4"
)

var allowedIPs = map[string]bool{
	"179.33.219.64": true,
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !allowedIPs[ip] {
		http.Redirect(w, r, "https://distriramirez.com.co/biometrico", http.StatusSeeOther)
		return
	}

	w.Write([]byte("¡Bienvenido!"))
}

func main() {

	e := echo.New()
	router.EchoRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

	// PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	// HOST := os.Getenv("SERVER_HOST")

	// e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", HOST, PORT)))
}
