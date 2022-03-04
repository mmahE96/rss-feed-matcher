package template

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

type hand int

func (h hand) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	//from form extract searchTerm string and deliver it to the function rss feed search
	//after searching, crate new struct and populate it with data, content and title
	//show fetched data on html template
	//repeat everything with angular and react
	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}

	//after adding request form parsing add response header setting
	//with header we can send response from server about way of interpretation response data in browser
	//data can be shown as plain text and as html just by changing header field

	tpl.ExecuteTemplate(w, "index.html", data)
	//fmt.Println(data)

	c := data.Submissions.Get("searchTerm")

	fmt.Println("only req.Form:", reflect.TypeOf(c), c)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func index() {
	var h hand
	http.ListenAndServe(":8080", h)
}
