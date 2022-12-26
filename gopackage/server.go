package main

import ( 
	"log"
	"net/http"
	"github.com/jetsadawwts/echoapi/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	Status  string `json:"status"`
}


var messages = []Message{
	{Status: "OK"},
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, messages)
}

func main() {
	user.InitDB()

	e := echo.New()
	e.Use(middleware.BasicAuth(func (username, password string, c echo.Context) (bool, error) {
		if username == "apidesign" && password == "45678" {
			return true, nil
		}
		return false, nil		
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)
	e.GET("/users", user.GetUsersHandler)
	e.POST("/users", user.CreateUsersHandler)
	e.GET("/users/:id", user.GetUserHandler)


	log.Fatal(e.Start(":2565"))
}
