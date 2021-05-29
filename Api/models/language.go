package models

import (
	"fmt"
	"net/http"

	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Language struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name" validate:"required"`
}

func NewLanguage() *Language {
	return &Language{}
}

func (l *Language) Url() string {
	return "/languages"
}

func (l *Language) DetailUrl() string {
	return "/languages/:id"
}

func (l *Language) GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var languages []Language
		db.Find(&languages)
		return c.JSON(http.StatusOK, languages)
	}

}

func (l *Language) Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		language := new(Language)

		if err := c.Bind(language); err != nil {
			return fmt.Errorf("id is mandatory for update")
		}

		if err := v.Validate(language); err != nil {
			return err
		}

		db.Save(language)

		return c.JSON(http.StatusCreated, language)
	}

}

func (l *Language) Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		lang := new(Language)
		if err := c.Bind(lang); err != nil {
			return err
		}

		langId := c.Param("id")

		if langId == "0" {
			return v.InvalidId()
		}

		if err := v.Validate(lang); err != nil {
			return err
		}

		db.Model(&lang).Where("id=?", langId).Updates(&lang)

		return c.JSON(http.StatusOK, lang)
	}

}

func (l *Language) Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error {
	return func(c echo.Context) error {
		langId := c.Param("id")
		lang := &Language{}
		db.First(&lang, langId)

		if langId == "0" {
			return v.InvalidId()
		}

		if lang != nil {
			db.Delete(&lang)
		}

		return c.JSON(http.StatusNoContent, lang)
	}
}
