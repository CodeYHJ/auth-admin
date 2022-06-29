package main

import (
	"github.com/codeyhj/auth-admin/server/db"
	"github.com/codeyhj/auth-admin/server/router"
	"github.com/codeyhj/auth-admin/server/util"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	e := echo.New()
	// validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()
	router.AddRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
