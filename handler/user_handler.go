package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// @Summary Get a user by ID
// @Produce json
// @Param id path int true "User ID" Example(1)
// @Success 200 {object} User
// @Router /user/{id} [get]
func getUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// TODO: Retrieve user from database or other data source based on ID

	user := User{
		ID:   id,
		Name: "Example User",
	}

	return c.JSON(http.StatusOK, user)
}

// @Summary Get a user by ID and name
// @Produce json
// @Param id query int true "User ID" Example(123)
// @Param name query string true "User Name" Example(jjj)
// @Success 200 {object} User
// @Router /user [get]
func getUserByIDAndName(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	name := c.QueryParam("name")

	// TODO: Retrieve user from database or other data source based on ID and name

	user := User{
		ID:   id,
		Name: name,
	}

	return c.JSON(http.StatusOK, user)
}
