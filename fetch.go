// GET the HTML content of a URL as an UTF-8 string
package fetch

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// contains normalized url and the response body string
type Page struct {
	Url  string
	Body string
}

// Perform a HTTP GET request, return struct with UTF-8 body
func Get(url string) (Page, error) {
	p := Page{Url: normalizeURL(url)}
	resp, err := http.Get(p.Url)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()
	str, err := responseToString(resp)
	if err != nil {
		return p, err
	}
	p.Body = str
	return p, nil
}

func normalizeURL(url string) string {
	switch {
	case regexp.MustCompile(`(?i)^https?`).MatchString(url):
		return url
	case regexp.MustCompile(`^//`).MatchString(url):
		return "http:" + url
	case regexp.MustCompile(`^/`).MatchString(url):
		return "http:/" + url
	default:
		return "http://" + url
	}
}

func responseToString(r *http.Response) (string, error) {
	// TODO: introspect http headers for encoding
	// TODO: handle non-utf8-text
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
