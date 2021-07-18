package player

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"github.com/anaskhan96/soup"
)

func GetClientID() (id string, err error) {
	respSoup, err := soup.Get("https://soundcloud.com/discover")
	if err != nil {
		return "null", errors.New("Error01")
	}
	doc := soup.HTMLParse(respSoup)
	scripts := doc.FindAll("script")
	var link string
	for i, script := range scripts {
		if i == (len(scripts) - 2) {
			link = script.Attrs()["src"]
		}
	}
	resp, err := http.Get(link)
	if err != nil {
		return "null", errors.New("Error02")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "null", errors.New("Error03")
	}
	r, err := regexp.Compile("client_id")
	if err != nil {
		return "null", errors.New("Error04")
	}
	k := r.FindStringIndex(string(body))[1]
	s := string(body)[k+2 : k+34]
	return s, nil
}
