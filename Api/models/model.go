package models

import (
	"net/http"

	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Model interface {
	GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	GetById(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error
	Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error

	Url() string
	DetailUrl() string
}

type GenericModel struct {
}

func NewGenericModel() *GenericModel {
	return &GenericModel{}
}

func (gm *GenericModel) GetFolderAndLanguageOptions(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {

		var languages []Language
		db.Find(&languages)

		var folders []Folder
		db.Find(&folders)

		fl := NewFolderAndLanguages(folders, languages)

		return c.JSON(http.StatusOK, fl)
	}
}
