package main

import (
	"os"

	"gitlab.com/GO-test/database"
	"gitlab.com/GO-test/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetupHandler for grouping API
func SetupHandler(e *echo.Echo) {
	db, err := database.GetConnection()
	if err != nil || db.Error != nil {
		nfLog.Fatal("Database connection: " + err.Error())
	}

	repository := user.CreateRepository(db)
	mainAPI := e.Group("/:provider/api/v" + os.Getenv("VERSION"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `method=${method}, uri=${uri}, status=${status},error=${error} \n`,
	}))

	// USER
	userHandlers := &user.Handler{}
	userHandlers.Repository = &repository
	users := mainAPI.Group("/user")
	userHandlers.SetRoutes(users)
}
