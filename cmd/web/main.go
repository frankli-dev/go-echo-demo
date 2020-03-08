// package main sets up a web application using echo framework. The app imports data from a Data.gov dataset to
// postgresql database and displays the data using view templates.
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
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer implements Renderer interface
type TemplateRenderer struct {
	templates *template.Template
}

// Render func implementation of Renderer interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// setupTemplateRenderer configures all view templates
func setupTemplateRenderer() echo.Renderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}

// entry point
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
	e.Renderer = setupTemplateRenderer()
	e.Static("/static", "public/static")

	http.Handler(e, lister)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
