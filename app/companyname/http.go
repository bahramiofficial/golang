package companyname

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pejamn.ir/models"
	models "pejamn.ir/pers/models"
	"strconv"
)

type Http struct {
}

var logic *Logic

func NewHttp(e *echo.Group, lg *Logic) {
	g := e.Group("/companyname/")
	//e.GET("/companyname/:id/", read)
	g.GET(":id/", read)
	g.POST("", create)
	logic = lg
}

func read(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))
	cm, err := logic.read(models.PK(id))
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"companyname": cm,
	})
}

func create(context echo.Context) error {
	var (
		cm  models.CompanyName
		err error
	)
	err = context.Bind(&cm)
	if err != nil {
		return err
	}

	if err = logic.Create(&cm); err != nil {
		return err
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"companyname": cm,
	})
}
