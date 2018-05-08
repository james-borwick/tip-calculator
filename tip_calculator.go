package tipCalculator

import (
	"html/template"
	"net/http"
	"strconv"
)

var tip int
var totalBill int

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
	billIn := r.FormValue("billInput")
	percentageIn := r.FormValue("percentageInput")

	b, _ := strconv.ParseFloat(billIn, 64)
	p, _ := strconv.ParseFloat(percentageIn, 64)

	tip := (b / 100) * p
	totalBill := b + tip

	t := strconv.FormatFloat(tip, 'f', -1, 64)
	tb := strconv.FormatFloat(totalBill, 'f', -1, 64)

	tString := t
	tbString := tb

	d := struct {
		Tip       string
		BillTotal string
	}{
		Tip:       tString,
		BillTotal: tbString,
	}

	tpl.ExecuteTemplate(w, "processor.html", d)
}