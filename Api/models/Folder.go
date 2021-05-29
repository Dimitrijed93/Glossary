package models

import (
	"fmt"
	"net/http"

	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Folder struct {
	Id                    uint     `gorm:"primaryKey" json:"id"`
	Name                  string   `json:"name" validate:"required"`
	SourceLanguage        Language `json:"sourceLanguage" gorm:"association_foreignkey:Id" validate:"required"`
	SourceLanguageID      int
	DestinationLanguage   Language `json:"destinationLanguage" gorm:"association_foreignkey:Id" validate:"required"`
	DestinationLanguageID int
}

func NewFolder() *Folder {
	return &Folder{}
}

func (f *Folder) Url() string {
	return "/folders"
}

func (f *Folder) DetailUrl() string {
	return "/folders/:id"
}

func (f *Folder) GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var folders []Folder
		db.Find(&folders)
		return c.JSON(http.StatusOK, folders)
	}

}

func (f *Folder) Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		folder := new(Folder)

		if err := c.Bind(folder); err != nil {
			echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("uanble to bind folder"))
		}

		if err := v.Validate(folder); err != nil {
			return err
		}

		db.Save(folder)

		return c.JSON(http.StatusCreated, folder)
	}

}

func (f *Folder) Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		folder := new(Folder)
		if err := c.Bind(folder); err != nil {
			return err
		}

		folderId := c.Param("id")

		if folderId == "0" {
			return v.InvalidId()
		}

		if err := v.Validate(folder); err != nil {
			return err
		}

		db.Model(&folder).Where("id=?", folderId).Updates(&folder)

		return c.JSON(http.StatusOK, folder)
	}

}

func (f *Folder) Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error {
	return func(c echo.Context) error {
		folderId := c.Param("id")
		folder := &Folder{}
		db.First(&folder, folderId)

		if folderId == "0" {
			return v.InvalidId()
		}

		if folder != nil {
			db.Delete(&folder)
		}

		return c.JSON(http.StatusNoContent, folder)
	}
}
