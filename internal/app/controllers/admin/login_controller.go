package controller

import (
	"net/http"
	"time"

	models "github.com/CASCOLOCO/boreal.git/internal/app/models/admin"
	"github.com/CASCOLOCO/boreal.git/internal/config/mysql"
	entity "github.com/CASCOLOCO/boreal.git/internal/interfaces/entity/admin"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	admin := entity.Admin{}
	err := c.Bind(&admin)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user models.Admin

	if err := mysql.DB.Table("admin_user").Where("email = ? and password = ?", admin.Email, admin.Password).Scan(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Usuario o contraseña incorrectaa")
	}

	if user.Password != admin.Password {
		return echo.NewHTTPError(http.StatusNotFound, "Usuario o contraseña incorrecta")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, _ := token.SignedString([]byte("secret"))

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
