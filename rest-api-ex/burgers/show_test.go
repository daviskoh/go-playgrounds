package burgers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShow(t *testing.T) {
	// build url
	url := "http://localhost:3000/burgers/:id"
	// issue request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	Show(resp, req, nil)

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
		t.Errorf(fmt.Sprintf("body.patty does not equal beef"))
	}

	fmt.Printf("%+v", resp)
}
