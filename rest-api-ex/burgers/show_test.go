package burgers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShow(t *testing.T) {
	burgerId := "4"

	// build url
	url := "http://localhost:3000/burgers/" + burgerId
	// issue request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	params := httprouter.Params{
		{Key: "id", Value: burgerId},
	}
	Show(resp, req, params)

	// investigate using constants
	if resp.Code != 200 {
		t.Errorf(fmt.Sprintf("request to url: %s did not return 200", url))
	}

	// extract body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	burger := Burger{}

	// unmarshal body
	err = json.Unmarshal(body, &burger)
	if err != nil {
		t.Fatal(err)
	}

	// check body
	if burger.Patty != "beef" {
		t.Errorf(fmt.Sprintf("body.Patty does not equal beef"))
	}

	if burger.Id != 4 {
		t.Errorf(fmt.Sprintf("body.Id does not equal 4"))
	}

	fmt.Printf("%+v", resp)
}
