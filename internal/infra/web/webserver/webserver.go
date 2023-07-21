package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: serverPort,
	}
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	http.ListenAndServe(s.WebServerPort, s.Router)
}
