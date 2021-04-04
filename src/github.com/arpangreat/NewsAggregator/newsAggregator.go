package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	//string_body := string(bytes)

	//fmt.Println(string_body)
	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
