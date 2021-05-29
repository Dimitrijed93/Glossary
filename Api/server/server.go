package server

import (
	models "github.com/dimitrijed93/glossary/api/models"
	util "github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitServer(model models.Model, db *gorm.DB, e *echo.Echo, v *util.Validator) {

	e.GET(model.Url(), model.GetAll(db, v))

	e.POST(model.Url(), model.Create(db, v))

	e.DELETE(model.DetailUrl(), model.Delete(db, v))

	e.PUT(model.DetailUrl(), model.Update(db, v))
}

func InitGenericModel(model *models.GenericModel, db *gorm.DB, e *echo.Echo, v *util.Validator) {

	e.GET("/folders-and-languages", model.GetFolderAndLanguageOptions(db, v))
}
