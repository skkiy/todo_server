package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/api/option"

	"github.com/sk62793/todo_server/config"
	"github.com/sk62793/todo_server/graph"
	"github.com/sk62793/todo_server/graph/generated"
	"github.com/sk62793/todo_server/repository"
)

const defaultPort = "8080"

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	tR := repository.NewRepository(db)

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: graph.NewResolver(tR)},
		),
	)
	ph := playground.Handler("GraphQL", "/query")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(auth)

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

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile("secret.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return err
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			return err
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return err
		}

		c.Set("token", token)
		return next(c)
	}
}
