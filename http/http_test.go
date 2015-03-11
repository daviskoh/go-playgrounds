package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEx(t *testing.T) {
	resp := httptest.NewRecorder()

	url := "http://localhost:3000/hello"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	Handler(resp, req)

	if resp.Code != 200 {
		t.Errorf(fmt.Sprintf("request to url: %s did not return 200", url))
	}
}
