package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olivere/elastic"
)

const esurl = "http://es01:9200"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/es", es)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func es(c echo.Context) error {

	client, err := elastic.NewClient(elastic.SetURL(esurl))
	if err != nil {
		// Handle error
	}
	fmt.Print("bb")
	info, code, err := client.Ping(esurl).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	m := "Elasticsearch returned with code " + strconv.Itoa(code) + " and version " + info.Version.Number + "\n"
	return c.String(http.StatusOK, m)
}
