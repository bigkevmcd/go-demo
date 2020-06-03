package demo

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// Config represents the configuration for the handler.
type Config struct {
	Redis *redis.Client
	Key   string
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
	pong, err := h.config.Redis.Get(r.Context(), h.config.Key).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "serving %s\n", pong)
}
