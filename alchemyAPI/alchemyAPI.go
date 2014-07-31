package alchemyAPI

import (
	"fmt"
	"net/url"
	"net/http"
	"os"
	"io"
	"io/ioutil"
	"encoding/json"
	"errors"
	"bytes"
)

type callMap map[string] map[string] string

type apiKey string

type Alchemist struct {
	m_key     apiKey
	m_callMap callMap
	m_baseURL string
}

func NewAlchemist(a_apiKey io.Reader) *Alchemist {
	buff := new(bytes.Buffer)
	buff.ReadFrom(a_apiKey)
	alchemist := new(Alchemist)
	alchemist.m_key = apiKey(buff.String())
	alchemist.m_callMap = callMap{
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
	alchemist.m_baseURL = "http://access.alchemyapi.com/calls"
	return alchemist
}

func (alchemist *Alchemist) Entities(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["entities"][flavor]
	if !ok {
		return nil, errors.New("Invalid Entities type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["entities"][flavor], options, nil)
}

func (alchemist *Alchemist) Keywords(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["keywords"][flavor]
	if !ok {
		return nil, errors.New("Invalid Keywords type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["keywords"][flavor], options, nil)
}

func (alchemist *Alchemist) Concepts(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["concepts"][flavor]
	if !ok {
		return nil, errors.New("Invalid Concepts type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["concepts"][flavor], options, nil)
}

func (alchemist *Alchemist) Sentiment(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["sentiment"][flavor]
	if !ok {
		return nil, errors.New("Invalid Sentiment type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["sentiment"][flavor], options, nil)
}

func (alchemist *Alchemist) TargetedSentiment(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["sentiment_targeted"][flavor]
	if !ok {
		return nil, errors.New("Invalid Targeted Sentiment type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["sentiment_targeted"][flavor], options, nil)
}

func (alchemist *Alchemist) Text(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["text"][flavor]
	if !ok {
		return nil, errors.New("Invalid text type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["text"][flavor], options, nil)
}

func (alchemist *Alchemist) TextRaw(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["text_raw"][flavor]
	if !ok {
		return nil, errors.New("Invalid text_raw type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["text_raw"][flavor], options, nil)
}

func (alchemist *Alchemist) Author(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["author"][flavor]
	if !ok {
		return nil, errors.New("Invalid author type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["author"][flavor], options, nil)
}

func (alchemist *Alchemist) Language(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["language"][flavor]
	if !ok {
		return nil, errors.New("Invalid language type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["language"][flavor], options, nil)
}

func (alchemist *Alchemist) Title(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["title"][flavor]
	if !ok {
		return nil, errors.New("Invalid title type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["title"][flavor], options, nil)
}

func (alchemist *Alchemist) Relations(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["relations"][flavor]
	if !ok {
		return nil, errors.New("Invalid relations type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["relations"][flavor], options, nil)
}

func (alchemist *Alchemist) Category(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["category"][flavor]
	if !ok {
		return nil, errors.New("Invalid category type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["category"][flavor], options, nil)
}

func (alchemist *Alchemist) Feeds(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["feeds"][flavor]
	if !ok {
		return nil, errors.New("Invalid feeds type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["feeds"][flavor], options, nil)
}

func (alchemist *Alchemist) Microformats(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["microformats"][flavor]
	if !ok {
		return nil, errors.New("Invalid microformats type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["microformats"][flavor], options, nil)
}

func (alchemist *Alchemist) Combined(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["combined"][flavor]
	if !ok {
		return nil, errors.New("Invalid combined type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["combined"][flavor], options, nil)
}

func (alchemist *Alchemist) Taxonomy(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["taxonomy"][flavor]
	if !ok {
		return nil, errors.New("Invalid taxonomy type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["taxonomy"][flavor], options, nil)
}

func (alchemist *Alchemist) ImageExtraction(flavor string, options url.Values, srcStr string) (map[string] interface{}, error) {
	_, ok := alchemist.m_callMap["image"][flavor]
	if !ok {
		return nil, errors.New("Invalid image extraction type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["image"][flavor], options, nil)
}

func (alchemist *Alchemist) ImageTagging(flavor string, options url.Values, imageStr string) (map[string] interface{}, error) {
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
	return  alchemist.analyze(alchemist.m_callMap["imagetagging"][flavor], options, imageFile)
}

func (alchemist *Alchemist) analyze(endpoint string, params url.Values, postData io.Reader) (map[string] interface{}, error) {
	params.Add("apikey", string(alchemist.m_key))
	params.Add("outputMode", "json")

	postURL := alchemist.m_baseURL + endpoint + "?" + params.Encode()
	
	var resp *http.Response
	var err error
	fmt.Println(postURL)
	if nil != postData {
		resp, err = http.Post(postURL, "image/jpeg", postData)
	} else {
		resp, err = http.Get(postURL)
	}
	contents, _ := ioutil.ReadAll(resp.Body)
	var jsonMap map[string] interface{}
	jsonErr := json.Unmarshal(contents, &jsonMap)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonMap, err
}
