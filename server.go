package main

import (
	"errors"
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

	tR, uR := repository.NewRepository(db)

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: graph.NewResolver(tR, uR)},
		),
	)
	ph := playground.Handler("GraphQL", "/query")

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.CORSAllowOrigin()},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// e.Use(auth)

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
		ctx := c.Request().Context()

		opt := option.WithCredentialsFile("/secret.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return errors.New("failed to initialize firebase.App")
		}
		client, err := app.Auth(ctx)
		if err != nil {
			return errors.New("failed to get auth client")
		}

		auth := c.Request().Header.Get("Authorization")

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errors.New("token is invalid")
		}

		if _, err := client.VerifyIDToken(ctx, parts[1]); err != nil {
			return errors.New("invalid token")
		}

		return next(c)
	}
}
