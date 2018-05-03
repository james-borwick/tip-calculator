package tipCalculator

import (
	"html/template"
    "net/http"
    "strconv"
)

// Page is a struct
// type Page struct {
//     NameOne string
// 	NameTwo string
// }

var tpl *template.Template

func init() {
    http.HandleFunc("/", index)
    http.HandleFunc("/results", processor)
    http.ListenAndServe(":8080", nil)
    tpl = template.Must(template.ParseGlob("*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
    billVarString := strconv.ParseInt (r.FormValue("billForm"), int, int)
    percentageVarString := r.FormValue("percentageForm")

    d := struct {
        Bill int
        Percentage int
    } {
        Bill: billVar,
        Percentage: percentageVar,
    }

    tpl.ExecuteTemplate(w, "results.html", d)
}
