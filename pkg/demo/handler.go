package demo

import (
	"fmt"
	"net/http"
)

// Config represents the configuration for the handler.
type Config struct {
	Name string
}

// Handler just serves HTTP requests.
type Handler struct {
	config Config
}

// New creates a new configuration
func New(c Config) *Handler {
	return &Handler{config: c}
}

// ServeHTTP implements the http.Handler interface.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "serving %s\n", h.config.Name)
}
