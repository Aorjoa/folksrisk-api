package main

import (
	"net/http"

	"github.com/aorjoa/folksrisk-api/handle"
	"github.com/aorjoa/folksrisk-api/model"
	"github.com/aorjoa/folksrisk-api/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Database
	db := service.Database()
	defer db.Close()

	//Auto migrate
	db.AutoMigrate(&model.Personal{})
	db.AutoMigrate(&model.Point{})
	db.AutoMigrate(&model.Evidence{})
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.BankAccount{})
	db.AutoMigrate(&model.PhoneNumber{})

	//initial person
	initPerson := &model.Personal{
		Identity: "A001",
		Name:     "Test Man",
		Email:    "test@gmail.com",
		Point: model.Point{
			Risk:              -10,
			PersinalIDVerify:  true,
			BankAccountVerify: false,
			PhoneNumberVerify: false,
			EmailActivated:    true,
			Sponsed:           false,
		},
		Image:        "images",
		BankAccounts: []model.BankAccount{model.BankAccount{BankAccount: "1931123115"}},
		PhoneNumbers: []model.PhoneNumber{model.PhoneNumber{PhoneNumber: "0993838382"}},
		Evidences:    []model.Evidence{model.Evidence{Description: "Test Comment"}},
		Files:        []model.File{model.File{FileName: "test1"}, model.File{FileName: "test2"}}}

	if (db.First(&model.Personal{}, "identity = ? and name = ? ", initPerson.Identity, initPerson.Name).RecordNotFound()) {
		db.Create(initPerson)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello")
	})

	e.GET("/list", handle.ListPerson)
	e.PUT("/list", handle.UpdatePerson)

	e.Logger.Fatal(e.StartTLS(":8443", "certs/cert.pem", "certs/key.pem"))
}
