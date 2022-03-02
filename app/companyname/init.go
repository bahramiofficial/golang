package companyname

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Group, db *gorm.DB) {
	repo := NewRepo(db)
	logic := NewLogic(repo)
	NewHttp(e, logic)
}
