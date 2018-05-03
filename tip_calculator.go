package tipCalculator

import (
	"html/template"
	"net/http"
)

// Page is a struct
type Page struct {
	NameOne string
	NameTwo string
}

var tpl *template.Template

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
	tpl = template.Must(template.ParseGlob("*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	fname := r.FormValue("nameNumberOne")
	lname := r.FormValue("nameNumberTwo")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}

	tpl.ExecuteTemplate(w, "processor.html", d)
}
