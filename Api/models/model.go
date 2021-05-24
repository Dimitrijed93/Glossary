package models

import (
	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Model interface {
	GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error

	Url() string
	DetailUrl() string
}
