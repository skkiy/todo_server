package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sk62793/todo_server/config"
	"github.com/sk62793/todo_server/graph"
	"github.com/sk62793/todo_server/graph/generated"
)

const defaultPort = "8080"

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// uR, rR, aR, tR, cR := repository.NewRepository(db)

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
	ph := playground.Handler("GraphQL", "/query")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/query", func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		ph.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}
