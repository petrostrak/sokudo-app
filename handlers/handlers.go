package handlers

import (
	"net/http"

	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/data"
)

// Handlers is the type for handlers, and gives access to Celeritas and models
type Handlers struct {
	App    *sokudo.Sokudo
	Models data.Models
}

// Home is the handler to render the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
