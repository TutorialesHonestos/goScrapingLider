package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Payload struct {
	Categories string `json:"categories"`
	Page       int    `json:"page"`
	//Facets      []interface{} `json:"facets"`
	Facets      []string `json:"facets"`
	Sortby      string   `json:"sortBy"`
	Hitsperpage int      `json:"hitsPerPage"`
}

type LiderScraping struct {
	Body Payload
}

func (l *LiderScraping) Init() {

	data := Payload{}

	data.Categories = "Computación/Computadores/Notebooks"
	data.Page = 1
	data.Facets = []string{}
	data.Sortby = ""
	data.Hitsperpage = 900

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://buysmart-bff-production.lider.cl/buysmart-bff/category", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "buysmart-bff-production.lider.cl")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("X-Flowid", "e545eeab-6404-4d6e-a370-f9de181be6a9")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://www.lider.cl")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.lider.cl/")
	req.Header.Set("Accept-Language", "es-ES,es;q=0.9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	bodyResponse, _ := ioutil.ReadAll(resp.Body)
	jsonResponse := string(bodyResponse)

	responseLider := ResponseLider{}
	decoder := json.NewDecoder(strings.NewReader(string(jsonResponse)))
	decoder.Decode(&responseLider)

	for _, producto := range responseLider.Products {
		fmt.Println(producto.Sku, producto.DisplayName, producto.Price.BasePriceReference)
	}

	defer resp.Body.Close()
}
