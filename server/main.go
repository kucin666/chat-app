package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thxrhmn/chat-app/database"
	"github.com/thxrhmn/chat-app/pkg/mysql"
	"github.com/thxrhmn/chat-app/routes"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()
	database.RunMigration()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	routes.RouteInit(e.Group("/api/v1"))

	var PORT = os.Getenv("PORT")

	fmt.Println("Server runing on localhost:" + PORT)

	e.Logger.Fatal(e.Start(":" + PORT))
}
