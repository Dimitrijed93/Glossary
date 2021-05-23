package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Type struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Value string `json:"value"`
}

func NewType() *Type {
	return &Type{}
}

func (t *Type) GetAll(db *gorm.DB) func(c echo.Context) error {

	return func(c echo.Context) error {
		var types []Type
		db.Find(&types)
		return c.JSON(http.StatusOK, types)
	}

}

func (t *Type) Create(db *gorm.DB) func(c echo.Context) error {

	return func(c echo.Context) error {
		typeObj := new(Type)
		if err := c.Bind(typeObj); err != nil {
			return err
		}
		db.Save(typeObj)

		return c.JSON(http.StatusCreated, typeObj)
	}

}

func (t *Type) Update(db *gorm.DB) func(c echo.Context) error {

	return func(c echo.Context) error {
		typeObj := new(Type)
		if err := c.Bind(typeObj); err != nil {
			return err
		}

		typeId := c.Param("id")

		db.Model(&typeObj).Where("id = ?", typeId).Updates(&typeObj)

		return c.JSON(http.StatusCreated, typeObj)
	}

}

func (t *Type) Delete(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		typeId := c.Param("id")
		typeObj := &Type{}
		db.First(&typeObj, typeId)

		if typeObj != nil {
			db.Delete(&typeObj)
		}

		return c.JSON(http.StatusNoContent, typeObj)
	}
}
