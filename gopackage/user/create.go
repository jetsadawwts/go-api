package user 


import (
	"net/http"
	"github.com/labstack/echo/v4"
)


func CreateUsersHandler(c echo.Context) error {
	u := User{}
	err := c.Bind(&u)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	raw :=  db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", u.Name, u.Age)
	err = raw.Scan(&u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	} 

	return c.JSON(http.StatusCreated, u)

}