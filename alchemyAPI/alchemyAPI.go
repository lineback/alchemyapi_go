package alchemyAPI

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type callMap map[string]map[string]string

type apiKey string

type Alchemist struct {
	m_key     apiKey
	m_callMap callMap
	m_baseURL string
}

/*
Creates new Alchemist
INPUT:
    Reader which contains API Key
*/
func NewAlchemist(a_apiKey io.Reader) *Alchemist {
	buff := new(bytes.Buffer)
	buff.ReadFrom(a_apiKey)
	alchemist := new(Alchemist)
	alchemist.m_key = apiKey(buff.String())
	alchemist.m_callMap = callMap{
		"sentiment": map[string]string{
			"url":  "/url/URLGetTextSentiment",
			"text": "/text/TextGetTextSentiment",
			"html": "/html/HTMLGetTextSentiment",
		},
		"sentiment_targeted": map[string]string{
			"url":  "/url/URLGetTargetedSentiment",
			"text": "/text/TextGetTargetedSentiment",
			"html": "/html/HTMLGetTargetedSentiment",
		},
		"author": map[string]string{
			"url":  "/url/URLGetAuthor",
			"html": "/html/HTMLGetAuthor",
		},
		"keywords": map[string]string{
			"url":  "/url/URLGetRankedKeywords",
			"text": "/text/TextGetRankedKeywords",
			"html": "/html/HTMLGetRankedKeywords",
		},
		"concepts": map[string]string{
			"url":  "/url/URLGetRankedConcepts",
			"text": "/text/TextGetRankedConcepts",
			"html": "/html/HTMLGetRankedConcepts",
		},
		"entities": map[string]string{
			"url":  "/url/URLGetRankedNamedEntities",
			"text": "/text/TextGetRankedNamedEntities",
			"html": "/html/HTMLGetRankedNamedEntities",
		},
		"category": map[string]string{
			"url":  "/url/URLGetCategory",
			"text": "/text/TextGetCategory",
			"html": "/html/HTMLGetCategory",
		},
		"relations": map[string]string{
			"url":  "/url/URLGetRelations",
			"text": "/text/TextGetRelations",
			"html": "/html/HTMLGetRelations",
		},
		"text": map[string]string{
			"url":  "/url/URLGetText",
			"html": "/html/HTMLGetText",
		},
		"text_raw": map[string]string{
			"url":  "/url/URLGetRawText",
			"html": "/html/HTMLGetRawText",
		},
		"title": map[string]string{
			"url":  "/url/URLGetTitle",
			"html": "/html/HTMLGetTitle",
		},
		"language": map[string]string{
			"url":  "/url/URLGetLanguage",
			"html": "/html/HTMLGetLanguage",
			"text": "/text/TextGetLanguage",
		},
		"feeds": map[string]string{
			"url":  "/url/URLGetFeedLinks",
			"html": "/html/HTMLGetFeedLinks",
		},
		"microformats": map[string]string{
			"url":  "/url/URLGetMicroformatData",
			"html": "/html/HTMLGetMicroformatData",
		},
		"combined": map[string]string{
			"url":  "/url/URLGetCombinedData",
			"text": "/text/TextGetCombinedData",
		},
		"image": map[string]string{
			"url": "/url/URLGetImage",
		},
		"imagetagging": map[string]string{
			"url":   "/url/URLGetRankedImageKeywords",
			"image": "/image/ImageGetRankedImageKeywords",
		},
		"taxonomy": map[string]string{
			"url":  "/url/URLGetRankedTaxonomy",
			"html": "/html/HTMLGetRankedTaxonomy",
			"text": "/text/TextGetRankedTaxonomy",
		},
	}
	alchemist.m_baseURL = "http://access.alchemyapi.com/calls"
	return alchemist
}

/*
Extracts the entities for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/entity-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/entity-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    disambiguate -> disambiguate entities (i.e. Apple the company vs. apple the fruit). 0: disabled, 1: enabled (default)
    linkedData -> include linked data on disambiguated entities. 0: disabled, 1: enabled (default)
    coreference -> resolve coreferences (i.e. the pronouns that correspond to named entities). 0: disabled, 1: enabled (default)
    quotations -> extract quotations by entities. 0: disabled (default), 1: enabled.
    sentiment -> analyze sentiment for each entity. 0: disabled (default), 1: enabled. Requires 1 additional API transction if enabled.
    showSourceText -> 0: disabled (default), 1: enabled
    maxRetrieve -> the maximum number of entities to retrieve (default: 50)

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Entities(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["entities"][flavor]
	if !ok {
		return nil, errors.New("Invalid Entities type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["entities"][flavor], options, nil)
}

/*
Extracts the keywords from text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/keyword-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/keyword-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    keywordExtractMode -> normal (default), strict
    sentiment -> analyze sentiment for each keyword. 0: disabled (default), 1: enabled. Requires 1 additional API transaction if enabled.
    showSourceText -> 0: disabled (default), 1: enabled.
    maxRetrieve -> the max number of keywords returned (default: 50)

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Keywords(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["keywords"][flavor]
	if !ok {
		return nil, errors.New("Invalid Keywords type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["keywords"][flavor], options, nil)
}

/*
Tags the concepts for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/concept-tagging/
    For the docs, please refer to: http://www.alchemyapi.com/api/concept-tagging/

Available Options:
    maxRetrieve -> the maximum number of concepts to retrieve (default: 8)
    linkedData -> include linked data, 0: disabled, 1: enabled (default)
    showSourceText -> 0:disabled (default), 1: enabled

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Concepts(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["concepts"][flavor]
	if !ok {
		return nil, errors.New("Invalid Concepts type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["concepts"][flavor], options, nil)
}

/*
Calculates the sentiment for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/sentiment-analysis/
    For the docs, please refer to: http://www.alchemyapi.com/api/sentiment-analysis/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    showSourceText -> 0: disabled (default), 1: enabled

OUTPUT:
    The response,converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Sentiment(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["sentiment"][flavor]
	if !ok {
		return nil, errors.New("Invalid Sentiment type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["sentiment"][flavor], options, nil)
}

/*
Calculates the targeted sentiment for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/sentiment-analysis/
    For the docs, please refer to: http://www.alchemyapi.com/api/sentiment-analysis/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    target -> the word or phrase to run sentiment analysis on.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    showSourceText-> 0: disabled, 1: enabled

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) TargetedSentiment(flavor string, options url.Values, srcStr string, target string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["sentiment_targeted"][flavor]
	if !ok {
		return nil, errors.New("Invalid Targeted Sentiment type!")
	}
	options.Add(flavor, srcStr)
	options.Add("target", target)
	return alchemist.analyze(alchemist.m_callMap["sentiment_targeted"][flavor], options, nil)
}

/*
Extracts the cleaned text (removes ads, navigation, etc.) for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/text-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/text-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    useMetadata -> utilize meta description data, 0: disabled, 1: enabled (default)
    extractLinks -> include links, 0: disabled (default), 1: enabled.

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Text(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["text"][flavor]
	if !ok {
		return nil, errors.New("Invalid text type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["text"][flavor], options, nil)
}

/*
Extracts the raw text (includes ads, navigation, etc.) for a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/text-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/text-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    none

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) TextRaw(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["text_raw"][flavor]
	if !ok {
		return nil, errors.New("Invalid text_raw type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["text_raw"][flavor], options, nil)
}

/*
Extracts the author from a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/author-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/author-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Availble Options:
    none

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Author(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["author"][flavor]
	if !ok {
		return nil, errors.New("Invalid author type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["author"][flavor], options, nil)
}

/*
Detects the language for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/api/language-detection/
    For the docs, please refer to: http://www.alchemyapi.com/products/features/language-detection/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    none

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Language(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["language"][flavor]
	if !ok {
		return nil, errors.New("Invalid language type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["language"][flavor], options, nil)
}

/*
Extracts the title for a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/text-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/text-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    useMetadata -> utilize title info embedded in meta data, 0: disabled, 1: enabled (default)

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Title(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["title"][flavor]
	if !ok {
		return nil, errors.New("Invalid title type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["title"][flavor], options, nil)
}

/*
Extracts the relations for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/relation-extraction/
    For the docs, please refer to: http://www.alchemyapi.com/api/relation-extraction/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    sentiment -> 0: disabled (default), 1: enabled. Requires one additional API transaction if enabled.
    keywords -> extract keywords from the subject and object. 0: disabled (default), 1: enabled.
                Requires one additional API transaction if enabled.
    entities -> extract entities from the subject and object. 0: disabled (default), 1: enabled.
                Requires one additional API transaction if enabled.
    requireEntities -> only extract relations that have entities. 0: disabled (default), 1: enabled.
    sentimentExcludeEntities -> exclude full entity name in sentiment analysis. 0: disabled, 1: enabled (default)
    disambiguate -> disambiguate entities (i.e. Apple the company vs. apple the fruit). 0: disabled, 1: enabled (default)
    linkedData -> include linked data with disambiguated entities. 0: disabled, 1: enabled (default).
    coreference -> resolve entity coreferences. 0: disabled, 1: enabled (default)
    showSourceText -> 0: disabled (default), 1: enabled.
    maxRetrieve -> the maximum number of relations to extract (default: 50, max: 100)

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Relations(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["relations"][flavor]
	if !ok {
		return nil, errors.New("Invalid relations type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["relations"][flavor], options, nil)
}

/*
Categorizes the text for text, a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/text-categorization/
    For the docs, please refer to: http://www.alchemyapi.com/api/text-categorization/

INPUT:
    flavor -> which version of the call, i.e. text, url or html.
    data -> the data to analyze, either the text, the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    showSourceText -> 0: disabled (default), 1: enabled

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Category(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["category"][flavor]
	if !ok {
		return nil, errors.New("Invalid category type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["category"][flavor], options, nil)
}

/*
Detects the RSS/ATOM feeds for a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/feed-detection/
    For the docs, please refer to: http://www.alchemyapi.com/api/feed-detection/

INPUT:
    flavor -> which version of the call, i.e.  url or html.
    data -> the data to analyze, either the the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    none

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Feeds(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["feeds"][flavor]
	if !ok {
		return nil, errors.New("Invalid feeds type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["feeds"][flavor], options, nil)
}

/*
Parses the microformats for a URL or HTML.
    For an overview, please refer to: http://www.alchemyapi.com/products/features/microformats-parsing/
    For the docs, please refer to: http://www.alchemyapi.com/api/microformats-parsing/

INPUT:
    flavor -> which version of the call, i.e.  url or html.
    data -> the data to analyze, either the the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    none

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Microformats(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["microformats"][flavor]
	if !ok {
		return nil, errors.New("Invalid microformats type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["microformats"][flavor], options, nil)
}

/*
Combined call for page-image, entity, keyword, title, author, taxonomy,  concept.

INPUT:
    flavor -> which version of the call, i.e.  url or html.
    data -> the data to analyze, either the the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.

Available Options:
    extract ->
        Possible values: page-image, entity, keyword, title, author, taxonomy,  concept
        default        : entity, keyword, taxonomy,  concept

    disambiguate ->
        disambiguate detected entities
        Possible values:
            1 : enabled (default)
            0 : disabled

    linkedData ->
        include Linked Data content links with disambiguated entities
        Possible values :
            1 : enabled (default)
            0 : disabled

    coreference ->
        resolve he/she/etc coreferences into detected entities
        Possible values:
            1 : enabled (default)
            0 : disabled

    quotations ->
        enable quotations extraction
        Possible values:
            1 : enabled
            0 : disabled (default)

    sentiment ->
        enable entity-level sentiment analysis
        Possible values:
            1 : enabled
            0 : disabled (default)

    showSourceText ->
        include the original 'source text' the entities were extracted from within the API response
        Possible values:
            1 : enabled
            0 : disabled (default)

    maxRetrieve ->
        maximum number of named entities to extract
        default : 50

    baseUrl ->
        rel-tag output base http url

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Combined(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["combined"][flavor]
	if !ok {
		return nil, errors.New("Invalid combined type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["combined"][flavor], options, nil)
}

/*
Taxonomy classification operations.

INPUT:
    flavor -> which version of the call, i.e.  url or html.
    data -> the data to analyze, either the the url or html code.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.


Available Options:
    showSourceText  ->
        include the original 'source text' the taxonomy categories were extracted from within the API response
        Possible values:
            1 - enabled
            0 - disabled (default)

    sourceText ->
        where to obtain the text that will be processed by this API call.
        Possible values:
            cleaned_or_raw  : cleaning enabled, fallback to raw when cleaning produces no text (default)
            cleaned         : operate on 'cleaned' web page text (web page cleaning enabled)
            raw             : operate on raw web page text (web page cleaning disabled)
            cquery          : operate on the results of a visual constraints query
                              Note: The 'cquery' http argument must also be set to a valid visual constraints query.
            xpath           : operate on the results of an XPath query
                              Note: The 'xpath' http argument must also be set to a valid XPath query.

    cquery ->
        a visual constraints query to apply to the web page.
    xpath ->
        an XPath query to apply to the web page.

    baseUrl ->
        rel-tag output base http url (must be uri-argument encoded)

OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) Taxonomy(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["taxonomy"][flavor]
	if !ok {
		return nil, errors.New("Invalid taxonomy type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["taxonomy"][flavor], options, nil)
}

/*
Extracts main image from a URL

INPUT:
    flavor -> which version of the call (url only currently).
    data -> URL to analyze
    options -> various parameters that can be used to adjust how the API works,
               see below for more info on the available options.

Available Options:
    extractMode ->
        trust-metadata  :  (less CPU intensive, less accurate)
        always-infer    :  (more CPU intensive, more accurate)
OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) ImageExtraction(flavor string, options url.Values, srcStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["image"][flavor]
	if !ok {
		return nil, errors.New("Invalid image extraction type!")
	}
	options.Add(flavor, srcStr)
	return alchemist.analyze(alchemist.m_callMap["image"][flavor], options, nil)
}

/*
Tags an input image
INPUT:
    flavor -> which version of the call only url or image.
    data -> the data to analyze, either the the url or path to image.
    options -> various parameters that can be used to adjust how the API works, see below for more info on the available options.
OUTPUT:
    The response, converted from JSON to a map[string] interface{}
*/
func (alchemist *Alchemist) ImageTagging(flavor string, options url.Values, imageStr string) (map[string]interface{}, error) {
	_, ok := alchemist.m_callMap["imagetagging"][flavor]
	if !ok {
		return nil, errors.New("Invalid image tagging type!")
	}
	var imageReader io.Reader = nil
	var imageBytes []byte
	var err error
	if "image" == flavor {
		imageBytes, err = ioutil.ReadFile(imageStr)
		if nil != err {
			return nil, err
		}
		imageReader = bytes.NewReader(imageBytes)
		options.Add("imagePostMode", "raw")
	} else {
		options.Add(flavor, imageStr)
	}
	return alchemist.analyze(alchemist.m_callMap["imagetagging"][flavor], options, imageReader)
}

func (alchemist *Alchemist) analyze(endpoint string, params url.Values, postData io.Reader) (map[string]interface{}, error) {
	params.Add("apikey", string(alchemist.m_key))
	params.Add("outputMode", "json")

	postURL := alchemist.m_baseURL + endpoint + "?" + params.Encode()

	var resp *http.Response
	var err error = nil
	if nil != postData {
		resp, err = http.Post(postURL, "image/jpeg", postData)
	} else {
		resp, err = http.Get(postURL)
	}
	contents, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	var jsonMap map[string]interface{}
	jsonErr := json.Unmarshal(contents, &jsonMap)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonMap, err
}
