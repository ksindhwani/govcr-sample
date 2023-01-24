package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ksindhwani/govcr-sample/pkg/handler"
	"github.com/seborama/govcr/v13"
)

type Response struct {
	/** Fill in the fields for the response struct which you want to fetch from API response
	For example
	Email string `json:"email"`

	**/
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", handler.ProductHandler).Methods("GET")

	vcr := govcr.NewVCR(
		govcr.NewCassetteLoader("MyCassette1.json"),
		govcr.WithRequestMatchers(govcr.NewMethodURLRequestMatchers()...),
	)

	url := "<Api url>"
	req, err := http.NewRequest("<Request Method>", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.SetBasicAuth("<basic auth username>", "<basic auth password>")

	resp, err := vcr.HTTPClient().Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var response Response
	if err := json.Unmarshal(bodyText, &response); err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	http.ListenAndServe(":9090", r)
}
