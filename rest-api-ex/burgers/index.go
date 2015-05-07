package burgers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// dummy data for response
type Burger struct {
	patty      string
	condiments []string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	burgers := []Burger{
		Burger{patty: "beef"},
		Burger{patty: "veggie"},
	}

	if err := json.NewEncoder(w).Encode(burgers); err != nil {
		http.NotFound(w, r)
	}
}
