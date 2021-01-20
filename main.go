package main

import (
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

var nfLog *logrus.Entry
var loggerConfiguration middleware.LoggerConfig

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	///// LOAD ENV FILE
	//specify env file
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	// load env data
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

}
func main() {
	e := echo.New()

	e.GET("/info", func(c echo.Context) error {
		zone, _ := time.Now().Zone()
		local, _ := time.Now().In(time.Local).Zone()
		return c.JSON(http.StatusOK, echo.Map{
			"version":            os.Getenv("VERSION"),
			"Client-IP":          c.RealIP(),
			"host":               c.Request().Host,
			"time.Now()":         time.Now(),
			"time.Now().Local()": time.Now().Local(),
			"Time-Zone":          zone + " - " + local,
		})
	})

	SetupHandler(e)
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
