package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"

	_ "./matchers"
	"./search"
)

type hand int

var term string

type news struct {
	Number  int
	Field   string
	Content string
}

var News []news

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
	term = ""
}

func (h hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	for i, n := range search.TransferResults {
		News = append(News, news{Number: i, Content: n.Content, Field: n.Field})

	}
	if len(News) > 1 {
		fmt.Println("Whole news:", News[1].Content, News[1].Field, News[1].Number)
	}

	data := struct {
		Method        string
		Submissions   url.Values
		SearchResults []news
	}{
		r.Method,
		r.Form,
		News,
	}

	tpl.ExecuteTemplate(w, "index.html", data)

	c := data.Submissions.Get("searchTerm")
	re := data.SearchResults
	if len(re) > 1 {
		//fmt.Println("only req.Form:", c, re[1].Content)
	}

	term = c
	if term != "" {
		search.Run(term)
	}

}

var tpl *template.Template

func main() {
	if term != "" {
		search.Run(term)
	}

	fmt.Println("Selam from main", search.TransferResults)
	var h hand
	http.ListenAndServe(":3000", h)

}
