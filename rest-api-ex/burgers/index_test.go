package burgers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGET(t *testing.T) {
	resp := httptest.NewRecorder()

	url := "http://localhost:3000/burgers"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	Index(resp, req, nil)

	if resp.Code != 200 {
		t.Errorf(fmt.Sprintf("request to url: %s did not return 200", url))
	}

	burgers := []Burger{}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &burgers)

	if err != nil {
		t.Fatal(err)
	}

	if burgers == nil {
		t.Errorf(fmt.Sprintf("request body was nil"))
	}

	fmt.Printf("%+v", resp)
}
