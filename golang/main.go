package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olivere/elastic"
)

type Person struct {
	Id         string `json:"ID"`
	LINE       int    `json:"LINE"`
	Prefecture string `json:"Prefecture"`
	Volume     string `json:"Volume"`
	Number     string `json:"Number"`
	Year       int    `json:"Year"`
	Month      int    `json:"Month"`
	Day        int    `json:"Day"`
	Title      string `json:"Title"`
	Speaker    string `json:"Speaker"`
	Utterance  string `json:"Utterance"`
}

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
	info, code, err := client.Ping(esurl).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	m := "Elasticsearch returned with code " + strconv.Itoa(code) + " and version " + info.Version.Number + "\n"
	// client.CreateIndex("test")
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	bytes, err := ioutil.ReadFile("Pref13_tokyo.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	put1, err := client.Index().
		Index("test").
		Type("TTT").
		Id("1").
		BodyJson(persons).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	return c.String(http.StatusOK, m)
}
