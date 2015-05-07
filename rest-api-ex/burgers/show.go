package burgers

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// pretend to create something
	burger := Burger{Patty: "beef"}
	// return created object

	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	err := enc.Encode(burger)
	if err != nil {
		// tell client something went wrong
	}

	w.Write(buffer.Bytes())
}
