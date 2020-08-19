package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


//type Rates struct {
//	XMLName xml.Name `xml:"rates"`
//	Rates   []Item   `xml:"item"`
//}
//
//type Item []struct {
//	XMLName xml.Name `xml:"item"`
//	From    string   `xml:"from"`
//	To string `xml:"to"`
//	In string `xml:"in"`
//	Out string `xml:"out"`
//	Amount string `xml:"amount"`
//	Minamount string `xml:"minamount"`
//	Maxamount string `xml:"maxamount"`
//	Param string `xml:"param"`
//	City string `xml:"city"`
//}

type Rates struct {
	Items  []Item `xml:"item"`

}
type Item struct {
	From      string   `xml:"from"`
	To        string   `xml:"to"`
	In        string   `xml:"in"`
	Out       string   `xml:"out"`
	Amount    string   `xml:"amount"`
	Minamount string   `xml:"minamount"`
	Maxamount string   `xml:"maxamount"`
	Param     string   `xml:"param"`
	City      string   `xml:"city"`
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/stop", StopServer)
	router.HandleFunc("/courses", Courses)
	log.Fatal(http.ListenAndServe(":80", router))
}

func Courses(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("https://test.cryptohonest.ru/request-exportxml.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Rates
	xml.Unmarshal(bytes, &s)
	b, err := json.Marshal(s.Items)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write([]byte(b))
}

func StopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}


