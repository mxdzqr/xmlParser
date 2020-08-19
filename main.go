package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

type Rates struct {
	XMLName xml.Name `xml:"rates"`
	Rates   []Item   `xml:"item"`
}
type RatesJson struct {
	XMLName xml.Name `json:"rates"`
	Rates   []Item   `json:"item"`
}
type Item struct {
	XMLName xml.Name `xml:"item"`
	From    string   `xml:"from"`
	To string `xml:"to"`
	In string `xml:"in"`
	Out string `xml:"out"`
	Amount string `xml:"amount"`
	Minamount string `xml:"minamount"`
	Maxamount string `xml:"maxamount"`
	Param string `xml:"param"`
	City string `xml:"city"`
}
type ItemJson struct {
	XMLName xml.Name `json:"item"`
	From    string   `json:"from"`
	To string `json:"to"`
	In string `json:"in"`
	Out string `json:"out"`
	Amount string `json:"amount"`
	Minamount string `json:"minamount"`
	Maxamount string `json:"maxamount"`
	Param string `json:"param"`
	City string `json:"city"`
}

func StopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func Courses(w http.ResponseWriter, r *http.Request) {
	xmlValue, err := http.Get("https://test.cryptohonest.ru/request-exportxml.xml")

	if err != nil {
		fmt.Println(err)
	}

	defer xmlValue.Body.Close()
	fmt.Println("Success, was read link")
	body, _ := ioutil.ReadAll(xmlValue.Body)
	var rates Rates
	xml.Unmarshal(body, &rates)
	for i := 0; i < len(rates.Rates); i++ {
		fmt.Fprint(w, rates.Rates[i])
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/stop", StopServer)
	router.HandleFunc("/courses", Courses)
	http.Handle("/", router)

	fmt.Println("Server is start..")
	http.ListenAndServe(":80", nil)


}