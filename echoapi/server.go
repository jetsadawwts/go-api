package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Message struct {
	Status  string `json:"status"`
}


var messages = []Message{
	{Status: "OK"},
}


var users = []User{
	{ID: 1, Name: "jetsadawwts", Age: 26},
	{ID: 2, Name: "wongwit", Age: 25},
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, messages)
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

type Err struct {
	Message string `json:"message"`
}

func createUsersHandler(c echo.Context) error {
		var u User
		err := c.Bind(&u)
		
		if err != nil {
			return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
		}

		users = append(users, u)

		return c.JSON(http.StatusCreated, u)

}

func main() {
	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("/api")
	g.Use(middleware.BasicAuth(func (username, password string, c echo.Context) (bool, error) {
		if username == "apidesign" && password == "45678" {
			return true, nil
		}
		return false, nil		
	}))

	
	g.GET("/users", getUsersHandler)
	g.POST("/users", createUsersHandler)


	log.Println("Server started at :2565")
	log.Fatal(e.Start(":2565"))
	log.Println("bye bye!")
}
