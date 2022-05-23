package handlers

import (
	"net/http"
	"time"

	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/data"
)

type Handlers struct {
	App    *sokudo.Sokudo
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	if err := h.render(w, r, "home", nil, nil); err != nil {
		h.printError("error rendering:", err)
	}

}
