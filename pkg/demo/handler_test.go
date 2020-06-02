package demo

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	h := New(Config{Name: "just-a-test"})
	srv := httptest.NewServer(h)
	defer srv.Close()

	res, err := srv.Client().Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if string(body) != "serving just-a-test\n" {
		t.Fatalf("response got %s", body)
	}
}
