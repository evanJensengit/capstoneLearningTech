package main

//simple web server
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //this writes the body to a file and creates the file if it is not there, the "0600" is code for read write
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) //gets the body from the file with the name of "filename"
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil //returns page reference with title and body
}

//http.ResponseWriter handles assembles http servers response and by writing to it we send data to http client. http.Request is data struc that represtents client http request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //we print what is after "w," to the w and that becomes the server response
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi this is handler1 saying fuck off %s!", r.URL.Path[1:]) //we print what is after "w," to the w and that becomes the server response
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	//the "/" tells the http package to handle all requests to web root /" with the function "handler"

	http.HandleFunc("/", handler)       //HandleFunc is a method build into http imported thing
	http.HandleFunc("/balls", handler1) //HandleFunc is a method build into http imported thing
	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) //log.Fatal used to record what is returned from listenandserve.

	//the 8080 says "listen on this port"
}
