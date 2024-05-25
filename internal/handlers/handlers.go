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

type pageData struct {
	Stocks        []stock.Stock
	GoogleMapsKey string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	pageData := new(pageData)

	// Get stock data
	stockFile := "./assets/data/stocks.json"
	f, err := os.Open(stockFile)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &pageData.Stocks)
	if err != nil {
		panic(err)
	}

	// Get api key(s)
	keyFile := "./assets/keys/maps.key"
	f, err = os.Open(keyFile)
	if err != nil {
		panic(err)
	}
	data, err = io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	pageData.GoogleMapsKey = string(data)

	tempfile := "./assets/templates/index.html"
	tmpl, err := template.New(path.Base(tempfile)).ParseFiles(tempfile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		panic(err)
	}
}
