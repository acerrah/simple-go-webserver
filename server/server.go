package server

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Server is a struct that holds the address and the handler
type Server struct {
}

// new returns a new Server instance
func New() *Server {
	return &Server{}
}

// Handles the given html page
func (s *Server) HandlePage(page string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, page[12:], nil)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Starts the server by listening to the given address
// and handles the html files in the templates directory
func (s *Server) Start(Address string, pages []string) error {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	for _, page := range pages {
		if string(page[len(page)-10:]) == "index.html" {
			http.HandleFunc("/", s.HandlePage(page))
		} else if string(page[len(page)-5:]) == ".html" {
			log.Println(page)
			http.HandleFunc("/"+page[12:len(page)-5], s.HandlePage(page))
		}
	}
	log.Println("Server is starting at", Address)
	return http.ListenAndServe(Address, nil)
}
