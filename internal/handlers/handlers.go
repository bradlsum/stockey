package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/bradlsum/stockey/internal/stock"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	stockFile := "./assets/data/stocks.json"
	f, err := os.Open(stockFile)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	stocks := new([]stock.Stock)
	err = json.Unmarshal(data, stocks)
	if err != nil {
		panic(err)
	}

	tempfile := "./assets/templates/html.tmpl"
	tmpl, err := template.New(path.Base(tempfile)).ParseFiles(tempfile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, stocks)
	if err != nil {
		panic(err)
	}
}
