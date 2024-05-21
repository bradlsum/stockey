package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/bradlsum/stockey/internal/stock"
)

func main() {
	// Handle command-line
	url := "https://cpw.crestonedigital.com/fishing/api/v1/stocking"

	// Request data
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	stocks := []stock.Stock{}
	json.Unmarshal(data, &stocks)

	f, err := os.Create("./assets/data/stocks.json")
	if err != nil {
		panic(err)
	}

	prettyD, err := json.MarshalIndent(stocks, "", "    ")
	if err != nil {
		panic(err)
	}
	_, err = f.Write(prettyD)
	if err != nil {
		panic(err)
	}

	// sort.Slice(stocks, func(i, j int) bool {
	// 	t1, _ := time.Parse("2006-01-02", stocks[i].ReportDate)
	// 	t2, _ := time.Parse("2006-01-02", stocks[j].ReportDate)
	// 	return t2.Before(t1)
	// })

	// Parse data

	// TODO: Load into db

	// tempfile := "./assets/templates/html.tmpl"
	// tmpl, err := template.New(path.Base(tempfile)).ParseFiles(tempfile)
	// if err != nil {
	// 	panic(err)
	// }

	// f, err := os.Create("out.html")
	// if err != nil {
	// 	panic(err)
	// }

	// err = tmpl.Execute(f, stocks)
	// if err != nil {
	// 	panic(err)
	// }
}
