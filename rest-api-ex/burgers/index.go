package burgers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	burgers := []Burger{
		Burger{Patty: "beef"},
		Burger{Patty: "veggie"},
	}

	if err := json.NewEncoder(w).Encode(burgers); err != nil {
		http.NotFound(w, r)
	}
}
