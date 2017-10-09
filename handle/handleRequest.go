package handle

import (
	"net/http"

	"github.com/aorjoa/folksrisk-api/model"
	"github.com/aorjoa/folksrisk-api/service"
	"github.com/labstack/echo"
)

var (
	db = service.Database()
)

func ListPerson(c echo.Context) error {
	personal := []model.Personal{}
	db.Find(&personal)
	return c.JSON(http.StatusOK, personal)
}

func UpdatePerson(c echo.Context) error {
	u := &model.Personal{}
	if err := c.Bind(u); err != nil {
		return err
	}
	db.Save(&u)
	return c.JSON(http.StatusOK, u)
}
