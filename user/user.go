package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/GO-test/database"
)

// Handler is
type Handler struct {
	Repository *PostgreRepository
}

// SetRoutes -
func (h *Handler) SetRoutes(e *echo.Group) {

	e.GET("/s", test)
	e.POST("/create-table-user", createTableUser)
	e.POST("/create-table-profile", createTableProfile)
	e.GET("", h.getUser)
	e.POST("", insertUser)
	e.PATCH("", updateUser)
	e.DELETE("", deleteUser)

	profile := e.Group("/profile")

	profile.POST("/", createUserAndProfile)
}

func test(c echo.Context) error {
	return c.JSON(http.StatusOK, "USER GRANTED")
}

func createTableUser(c echo.Context) error {
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't connect to DB with error %v", err))
	}
	defer db.Close()

	users := new(User)

	users.CreateTable(db)
	return c.JSON(http.StatusOK, "User Table Created")
}

func createTableProfile(c echo.Context) error {
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't connect to DB with error %v", err))
	}
	defer db.Close()

	profiles := new(Profile)

	profiles.CreateTable(db)

	db.Model(&User{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")

	return c.JSON(http.StatusOK, "User Table Created")
}

func insertUser(c echo.Context) error {
	newUser := new(User)

	err := c.Bind(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't connect to DB with error %v", err))
	}

	rowsAffected := db.Where("email = ?", newUser.Email).First(&newUser).RowsAffected
	if rowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "Email sudah terdaftar")
	}
	response := db.Create(&newUser)
	if response.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal menambahkan user")
	}
	return c.JSON(http.StatusCreated, response.Value)
}

func (h *Handler) getUser(c echo.Context) error {
	email := c.QueryParam("email")
	users, err := h.Repository.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't connect to DB with error %v", err))
	}
	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	updatedUser := new(User)

	err := c.Bind(updatedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if updatedUser.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ID tidak dapat kosong")
	}
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't connect to DB with error %v", err))
	}
	response := db.Model(&updatedUser).Updates(updatedUser)
	if response.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "ID tidak ditemukan")
	}
	return c.JSON(http.StatusOK, response.Value)
}

func deleteUser(c echo.Context) error {
	deletedUser := new(User)

	err := c.Bind(deletedUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if deletedUser.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ID tidak dapat kosong")
	}
	db, err := database.GetConnection()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("can't connect to DB with error %v", err))
	}
	response := db.Delete(&deletedUser)
	if response.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "ID tidak ditemukan")
	}
	return c.JSON(http.StatusOK, response.Value)
}

func createUserAndProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
