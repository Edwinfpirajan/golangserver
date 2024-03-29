package controller

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/CASCOLOCO/boreal.git/internal/app/models"
	"github.com/CASCOLOCO/boreal.git/internal/config/mysql"
	"github.com/CASCOLOCO/boreal.git/internal/interfaces/entity"
	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	employeesWithSchedule := []models.EmployeeWithSchedule{}

	mysql.DB.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Scan(&employeesWithSchedule)
	return c.JSON(http.StatusOK, employeesWithSchedule)
}

func GetEmploye(c echo.Context) error {
	id := c.Param("id")

	employeeWithSchedule := models.EmployeeWithSchedule{}

	mysql.DB.Table("employes").Select("*").Joins("left join horaries h on h.id_sch = employes.schedule_id").Where("employes.id = ?", id).First(&employeeWithSchedule)

	if employeeWithSchedule.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Employee not found")
	}

	return c.JSON(http.StatusOK, employeeWithSchedule)
}

func SaveEmploye(c echo.Context) error {
	employe := entity.Employe{}

	err := c.Bind(&employe)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	employeFromDb := models.Employe{}

	mysql.DB.Table("employes").Where("employes.pin_employe = ?", employe.PinEmploye).Scan(&employeFromDb)

	if employeFromDb.ID == 0 {

		num1 := make([]byte, 1)
		rand.Read(num1)
		num1[0] = num1[0]%10 + 48

		num2 := make([]byte, 1)
		rand.Read(num2)
		num2[0] = num2[0]%10 + 48

		letter1 := make([]byte, 1)
		rand.Read(letter1)
		letter1[0] = letter1[0]%26 + 65

		letter2 := make([]byte, 1)
		rand.Read(letter2)
		letter2[0] = letter2[0]%26 + 65

		pin := fmt.Sprintf("%c%c%d%d", letter1[0], letter2[0], num1[0]-48, num2[0]-48)

		employeFromDb.PinEmploye = pin
		employeFromDb.FirstName = employe.FirstName
		employeFromDb.LastName = employe.LastName
		employeFromDb.Company = employe.Company
		employeFromDb.Position = employe.Position
		employeFromDb.ScheduleId = employe.ScheduleId

		mysql.DB.Save(&employeFromDb)

	} else {
		employeFromDb.FirstName = employe.FirstName
		employeFromDb.LastName = employe.LastName
		employeFromDb.Company = employe.Company
		employeFromDb.Position = employe.Position
		employeFromDb.ScheduleId = employe.ScheduleId

		mysql.DB.Save(&employeFromDb)
	}

	return c.JSON(http.StatusCreated, employe)
}

func DeleteEmploye(c echo.Context) error {
	id := c.Param("id")

	employee := models.Employe{}

	mysql.DB.Find(&employee, id)

	if employee.ID > 0 {
		mysql.DB.Delete(employee)
		return c.JSON(http.StatusOK, employee)
	} else {
		return echo.NewHTTPError(http.StatusNotFound, "Employee not found")
	}
}
