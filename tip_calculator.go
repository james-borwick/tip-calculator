package tipCalculator

import (
	"html/template"
	"net/http"
	"strconv"
)

var tip int
var totalBill int
var b int
var p int
var hundred int

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

	b := billIn
	p := percentageIn
	hundred := 100

	strconv.Atoi(b)
	strconv.Atoi(p)

	tip = (b / hundred) * p
	totalBill = b + tip

	d := struct {
		Tip       string
		BillTotal string
	}
	//{
		//Tip:       tip,
		//BillTotal: totalBill,
	//}

	tpl.ExecuteTemplate(w, "processor.html", d)
}
