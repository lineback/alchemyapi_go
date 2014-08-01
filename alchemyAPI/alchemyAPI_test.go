package alchemyAPI

import (
	"testing"
	"io/ioutil"
	"strings"
	"net/url"
)

func Test_Entities(t *testing.T) {
	keyBytes, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		t.Error("Entities failed. Error reading api key file")
	}
	alchemist := NewAlchemist(strings.NewReader(string(keyBytes[:40])))
	resp, respErr := alchemist.Entities("text", url.Values{}, "Bob broke my heart, and then made up this silly sentence to test the Go SDK")
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Enities FAILED")
	} else {
		t.Log("Text Entities PASSED")
	}
	resp, err = alchemist.Entities("html", url.Values{}, "<html><head><title>The best SDK Test | AlchemyAPI</title></head><body><h1>Hello World!</h1><p>My favorite language is PHP</p></body></html>")
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Enities FAILED")
	} else {
		t.Log("HTML Entities PASSED")
	}
	resp, err = alchemist.Entities("url", url.Values{}, "http://www.nytimes.com/2013/07/13/us/politics/a-day-of-friction-notable-even-for-a-fractious-congress.html?_r=0")
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Enities FAILED")
	} else {
		t.Log("URL Entities PASSED")
	}
}
