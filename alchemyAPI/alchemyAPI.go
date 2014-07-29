package alchemyAPI

import (
	"fmt"
	"net/url"
	"net/http"
	"os"
	"io"
)

type callMap map[string] map[string] string

type apiKey string

var g_apiKey apiKey
var g_callMap callMap
var g_baseURL string

func Init(a_apiKey ...string) {
	g_apiKey = apiKey(a_apiKey[0])
	g_callMap = callMap{
		"sentiment" : map[string]string { 
			"url"  : "/url/URLGetTextSentiment" ,
			"text" : "/text/TextGetTextSentiment",
			"html" : "/html/HTMLGetTextSentiment",
		},
		"sentiment_targeted" : map[string]string {
			"url"  : "/url/URLGetTextTargetedSentiment" ,
			"text" : "/text/TextGetTextTargetedSentiment",
			"html" : "/html/HTMLGetTextTargetedSentiment",
		},
		"author" : map[string]string {
			"url"  : "/url/URLGetAuthor" ,
			"html" : "/html/HTMLGetAuthor",
		},
		"keywords" : map[string]string {
			"url"  : "/url/URLGetRankedKeywords" ,
			"text" : "/text/TextGetRankedKeywords",
			"html" : "/html/HTMLGetRankedKeywords",
		},
		"concepts" : map[string]string {
			"url"  : "/url/URLGetRankedConcepts" ,
			"text" : "/text/TextGetRankedConcepts",
			"html" : "/html/HTMLGetRankedConcepts",
		},
		"entities" : map[string]string {
			"url"  : "/url/URLGetRankedNamedEntities" ,
			"text" : "/text/TextGetRankedNamedEntities",
			"html" : "/html/HTMLGetRankedNamedEntities",
		},
		"category" : map[string]string {
			"url"  : "/url/URLGetCategory" ,
			"text" : "/text/TextGetCategory",
			"html" : "/html/HTMLGetCategory",
		},
		"relations" : map[string]string {
			"url"  : "/url/URLGetRelations" ,
			"text" : "/text/TextGetRelations",
			"html" : "/html/HTMLGetRelations",
		},
		"text" : map[string]string {
			"url"  : "/url/URLGetText" ,
			"html" : "/html/HTMLGetText",
		},
		"title" : map[string]string {
			"url"  : "/url/URLGetTitle" ,
			"html" : "/html/HTMLGetTitle",
		},
		"feeds" : map[string]string {
			"url"  : "/url/URLGetFeedLinks" ,
			"html" : "/html/HTMLGetFeedLinks",
		},
		"microformats" : map[string]string {
			"url"  : "/url/URLGetMicroformatData" ,
			"html" : "/html/HTMLGetMicroformatData",
		},
		"combined" : map[string]string {
			"url"  : "/url/URLGetCombinedData" ,
			"html" : "/html/HTMLGetCombinedData",
		},
		"image" : map[string]string {
			"url"  : "/url/URLGetText" ,
		},
		"imagetagging" : map[string]string {
			"url"  : "/url/URLGetRankedImageKeywords" ,
			"image" : "/image/ImageGetRankedImageKeywords",
		},
		"taxonomy" : map[string]string {
			"url"  : "/url/URLGetRankedTaxonomy" ,
			"html" : "/html/HTMLGetRankedTaxonomy",
			"text" : "/text/TextGetRankedTaxonomy",
		},
	}
	g_baseURL = "http://access.alchemyapi.com/calls"
}

func Print() {
	fmt.Println(g_apiKey)
	fmt.Println(g_callMap)
}

func ImageTagging(flavor string, options url.Values, imageStr string) (*http.Response, error) {
	var imageFile io.Reader = nil
	var err error
	if "image" == flavor {
		imageFile, err = os.Open(imageStr)
		if nil != err {
			return nil, err
		}
	} else {
		options.Add(flavor, imageStr)
	}
	return  analyze(g_callMap["imagetagging"][flavor], options, imageFile)
}

func analyze(endpoint string, params url.Values, postData io.Reader) (*http.Response, error) {
	params.Add("apikey", string(g_apiKey))
	params.Add("outputMode", "json")

	postURL := g_baseURL + endpoint + "?" + params.Encode()
	
	var resp *http.Response
	var err error
	fmt.Println(postURL)
	if nil != postData {
		resp, err = http.Post(postURL, "image/jpeg", postData)
	} else {
		resp, err = http.Get(postURL)
	}
	return resp, err
}
