package main

import (
	"net/http"

	models "github.com/dimitrijed93/glossary/api/models"
	consts "github.com/dimitrijed93/glossary/api/util"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(postgres.Open(consts.CONNECTION_STRING), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Folder{})
	db.AutoMigrate(&models.ExtraItem{})
	db.AutoMigrate(&models.Type{})
	db.AutoMigrate(&models.EntryItem{})
	db.AutoMigrate(&models.Entry{})

	sqlDB, err := db.DB()

	if err != nil {
		panic("failed to get db object")
	}

	defer sqlDB.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	typeObj := models.NewType()

	e.GET(consts.TYPE_URL, typeObj.GetAll(db))

	e.POST(consts.TYPE_URL, typeObj.Create(db))

	e.DELETE(consts.TYPE_BY_ID_URL, typeObj.Delete(db))

	e.PUT(consts.TYPE_BY_ID_URL, typeObj.Update(db))

	e.Logger.Fatal(e.Start(consts.PORT))

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

	return db

}
