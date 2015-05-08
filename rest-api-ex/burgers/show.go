package burgers

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		// something went wrong w/ httprouter
	}

	// pretend to create something
	burger := Burger{Id: id, Patty: "beef"}
	// return created object

	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	err = enc.Encode(burger)
	if err != nil {
		// tell client something went wrong
	}

	w.Write(buffer.Bytes())
}
