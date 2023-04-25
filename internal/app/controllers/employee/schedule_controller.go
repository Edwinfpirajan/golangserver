package employee

import (
	"net/http"

	models "github.com/CASCOLOCO/boreal.git/internal/app/models/employee"
	"github.com/CASCOLOCO/boreal.git/internal/config/mysql"
	entity "github.com/CASCOLOCO/boreal.git/internal/interfaces/entity/employee"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllSchedule(c echo.Context) error {
	schedule := []models.Horary{}
	mysql.DB.Find(&schedule)
	return c.JSON(http.StatusOK, schedule)
}

//****************************************PARSER TIME

func saveSchedule(horary entity.Horary) (models.Horary, error) {

	schedule := models.Horary{}

	schedule.Id_sch = horary.Id_sch

	schedule.Arrival = horary.Arrival

	schedule.Departure = horary.Departure

	err := mysql.DB.Save(&schedule).Error
	if err != nil {
		return models.Horary{}, err
	}
	return schedule, nil
}

func SaveSchedule(c echo.Context) error {
	requestBody := entity.Horary{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	schedule, err := saveSchedule(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, schedule)
}

func DeleteSchedule(c echo.Context) error {
	id := entity.Horary{}

	if err := c.Bind(&id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var schedule models.Horary
	if err := mysql.DB.First(&schedule, id.Id_sch).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Schedule not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := mysql.DB.Delete(&schedule).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Schedule deleted")
}
