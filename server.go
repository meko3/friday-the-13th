package main

import (
	"net/http"
	"time"
	"html/template"
	"io"
	
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/", Home)
	e.Logger.Fatal(e.Start(":1323"))
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", template.HTML(genHTML(day())))
}

func day() []time.Time {
	n := time.Now()
    // start & end
    s := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
	e := time.Date(n.Year() + 10, n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
	days := fridayThe13th(s, e)
	return days
}

func fridayThe13th(s time.Time, e time.Time) []time.Time {
    array := []time.Time{}
	i := s
	for {
		if i.Day() == 13 && i.Weekday() == time.Friday {
			array = append(array, i)
		}
		if i == e {
			break
		}
 		i = i.AddDate(0, 0, 1)
	}
	return array
}

func genHTML(array []time.Time) string {
	html := ""
	for i := 0; i < len(array); i++ {
		if i + 1 < len(array) && array[i + 1].Year() == array[i].Year() {
			h := "<li>" + array[i].Format("2006-01-02") + "</li>"
			for {
				i++
				h += "<li>" + array[i].Format("2006-01-02") + "</li>"
				if i + 1 >= len(array) || array[i + 1].Year() != array[i].Year() {
					break
				}
			}
			html += "<li><ul class='year'>" + h + "</ul></li>";
		} else {
			html += "<li>" + array[i].Format("2006-01-02") + "</li>"
		}
	}
	return html
}
