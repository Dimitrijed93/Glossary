package main

import (
	"net/http"

	models "github.com/dimitrijed93/glossary/api/models"
	"github.com/dimitrijed93/glossary/api/server"
	util "github.com/dimitrijed93/glossary/api/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(postgres.Open(util.CONNECTION_STRING_LOCAL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Folder{})
	db.AutoMigrate(&models.ExtraItem{})
	db.AutoMigrate(&models.Type{})
	db.AutoMigrate(&models.EntryItem{})
	db.AutoMigrate(&models.Entry{})
	db.AutoMigrate(&models.Language{})

	sqlDB, err := db.DB()

	v := util.NewValidator()

	typeObj := models.NewType()
	folder := models.NewFolder()
	item := models.NewExtraItem()
	entry := models.NewEntry()
	entryItem := models.NewEntryItem()
	lang := models.NewLanguage()
	gm := models.NewGenericModel()

	server.InitServer(typeObj, db, e, v)
	server.InitServer(lang, db, e, v)
	server.InitServer(folder, db, e, v)
	server.InitServer(item, db, e, v)
	server.InitServer(entry, db, e, v)
	server.InitServer(entryItem, db, e, v)
	server.InitGenericModel(gm, db, e, v)

	if err != nil {
		panic("failed to get db object")
	}

	defer sqlDB.Close()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(util.PORT))

}

func ConnectToDB() *gorm.DB {

	dsn := "host=glossaryDB user=postgres password=dimit dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Belgrade"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Folder{})
	db.AutoMigrate(&models.ExtraItem{})
	db.AutoMigrate(&models.Type{})
	db.AutoMigrate(&models.EntryItem{})
	db.AutoMigrate(&models.Entry{})
	db.AutoMigrate(&models.Language{})

	return db

}
