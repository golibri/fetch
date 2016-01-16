package fetch

import (
	"testing"
)

func TestUrl(t *testing.T) {
	page, err := Get("http://blog.golang.org/go1.5")
	if err != nil {
		t.Fatal(err)
	}
	if page.Body == "" {
		t.FailNow()
	}
}

func TestIncompleteUrls(t *testing.T) {
	_, err1 := Get("//blog.golang.org/go1.5")
	if err1 != nil {
		t.FailNow()
	}
	_, err2 := Get("/blog.golang.org/go1.5")
	if err2 != nil {
		t.FailNow()
	}
	_, err3 := Get("blog.golang.org/go1.5")
	if err3 != nil {
		t.FailNow()
	}
}

func TestInvalidUrl(t *testing.T) {
	_, err := Get("invalid url")
	if err == nil {
		t.FailNow()
	}
}
