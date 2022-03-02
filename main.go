package main

import (
	"net/http"
	"pejamn.ir/app/companyname"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectionString() string {
	dsn := "host=localhost user=postgres password=12345 dbname=go port=5432 sslmode=disable TimeZone=Asia/Tehran"
	return dsn
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.CompanyName{}, &CarModel{}, &CreateYear{}, &Diversity{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := Connect()

	if err != nil {
		panic("server error ")
	}
	_ = db

	ech := echo.New()
	// ech.AutoTLSManager
	ech.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"healthcheck": true,
		})
	})

	grApi := ech.Group("/api/v1")
	companyname.New(grApi, db)

	ech.Logger.Fatal(ech.Start(":3000"))

}
