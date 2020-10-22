package controllers

import (
	"espadafajar/restku/helpers"
	"espadafajar/restku/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	usernam := c.FormValue("usernam")
	password := c.FormValue("password")

	res, err := models.CheckLogin(usernam, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	//generatetoken
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["usernam"] = usernam
	claims["level"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("uid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
