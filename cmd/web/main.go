// TODO: main comment
package main

import (
	"flag"
	"html/template"
	"io"
	"log"

	"github.com/ivan-marquez/golang-demo/pkg/http"
	"github.com/ivan-marquez/golang-demo/pkg/listing"
	"github.com/ivan-marquez/golang-demo/pkg/storage/pq"
	"github.com/ivan-marquez/golang-demo/scripts/data"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// TODO: add comment
type Template struct {
	templates *template.Template
}

// TODO: add comment
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// TODO: add comment
func setupTemplates() echo.Renderer {
	return &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}

func main() {
	var lister listing.Service
	storage, err := pq.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	migrate := flag.Bool("migrate", false, "migrate and populate database")
	flag.Parse()

	if *migrate {
		data.Migrate(storage.DB)
	}

	lister = listing.NewService(storage)

	// Echo config
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = setupTemplates()
	e.Static("/static", "public/static")

	http.Handler(e, lister)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
