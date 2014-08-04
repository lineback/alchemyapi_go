package main

import (
	"github.com/lineback/alchemyapi_go/alchemyAPI"
	"net/url"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func entitiesExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#   Entity Extraction Example              #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Entities("text", url.Values{"sentiment":{"1"}}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
 	fmt.Println("## Entitites ##")
	for _, keyword := range resp["entities"].([]interface{}){
		fmt.Println("text:", keyword.(map[string] interface{})["text"])
	 	fmt.Println("type:", keyword.(map[string] interface{})["type"])
		fmt.Println("relevance:", keyword.(map[string] interface{})["relevance"])
		fmt.Println("sentiment:", keyword.(map[string] interface{})["sentiment"].(map[string] interface{})["type"])
	}
}

func keywordsExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#   Keywords Extraction Example            #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Keywords("text", url.Values{"sentiment":{"1"}}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Keywords ##")
	for _, keyword := range resp["keywords"].([]interface{}){
		fmt.Println("text:", keyword.(map[string] interface{})["text"])
		fmt.Println("relevance:", keyword.(map[string] interface{})["relevance"])
		fmt.Println("sentiment:", keyword.(map[string] interface{})["sentiment"].(map[string] interface{})["type"])
	}
}

func conceptsExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#   Concepts Extraction Example            #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Concepts("text", url.Values{}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Concepts ##")
	for _, keyword := range resp["concepts"].([]interface{}){
		fmt.Println("text:", keyword.(map[string] interface{})["text"])
		fmt.Println("relevance:", keyword.(map[string] interface{})["relevance"])
	}
}

func sentimentExample(alchemist *alchemyAPI.Alchemist, demoHTML string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#   Sentiment Example                      #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing HTML: ", demoHTML)
	fmt.Println("")
	resp, err := alchemist.Sentiment("html", url.Values{}, demoHTML)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Document Sentiment ##")
	fmt.Println("type:", resp["docSentiment"].(map[string] interface{})["type"])
	fmt.Println("score:", resp["docSentiment"].(map[string] interface{})["score"])
}

func targetedSentimentExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#        Targeted Sentiment Example        #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.TargetedSentiment("html", url.Values{}, demoText, "Denver")
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Document Sentiment ##")
	fmt.Println("type:", resp["docSentiment"].(map[string] interface{})["type"])
	fmt.Println("score:", resp["docSentiment"].(map[string] interface{})["score"])
}

func textExtractionExample(alchemist *alchemyAPI.Alchemist, demoURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#        Text Extraction Example           #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing URL: ", demoURL)
	fmt.Println("")
	resp, err := alchemist.Text("url", url.Values{}, demoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Text ##")
	fmt.Println("Text:", resp["text"])
}

func authorExtractionExample(alchemist *alchemyAPI.Alchemist, demoURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#      Author Extraction Example           #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing URL: ", demoURL)
	fmt.Println("")
	resp, err := alchemist.Author("url", url.Values{}, demoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Author ##")
	fmt.Println("Author:", resp["author"])
}

func titleExtractionExample(alchemist *alchemyAPI.Alchemist, demoURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#      Title Extraction Example            #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing URL: ", demoURL)
	fmt.Println("")
	resp, err := alchemist.Title("url", url.Values{}, demoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Title ##")
	fmt.Println("Title:", resp["title"])
}

func languageDetectionExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#      Language Detection Example          #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Language("text", url.Values{}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Language ##")
	fmt.Println("Language:", resp["language"])
	fmt.Println("iso-693-1:", resp["iso-693-1"])
	fmt.Println("native speakers:", resp["native-speakers"])
}

func relationsExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#           Relations Example              #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Relations("text", url.Values{}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Relations ##")
	for _, relation := range resp["relations"].([]interface{}){
		fmt.Println("subject:", relation.(map[string] interface{})["subject"].(map [string] interface{})["text"])
		fmt.Println("action:", relation.(map[string] interface{})["action"].(map [string] interface{})["text"])
		fmt.Println("object:", relation.(map[string] interface{})["object"].(map [string] interface{})["text"])
	}
}

func textCategorizationExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#      Text Categorization Example         #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Category("text", url.Values{}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Category ##")
	fmt.Println("Category:", resp["category"])
	fmt.Println("Score:", resp["score"])
}

func feedDetectionExample(alchemist *alchemyAPI.Alchemist, demoURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#           Feeds Example                  #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing url: ", demoURL)
	fmt.Println("")
	resp, err := alchemist.Feeds("url", url.Values{}, demoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Feeds ##")
	for _, feed := range resp["feeds"].([]interface{}){
		fmt.Println("feed:", feed.(map[string] interface{})["feed"])
	}
}

func microformatsExample(alchemist *alchemyAPI.Alchemist, demoURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#      Microformats Example                #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing url: ", demoURL)
	fmt.Println("")
	resp, err := alchemist.Microformats("url", url.Values{}, demoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Microformats ##")
	for _, microformat := range resp["microformats"].([]interface{}){
		fmt.Println("field:", microformat.(map[string] interface{})["field"])
		fmt.Println("data:", microformat.(map[string] interface{})["data"])
	}
}

func imageExtractionExample(alchemist *alchemyAPI.Alchemist, demoURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#        Image Extraction Example          #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing url: ", demoURL)
	fmt.Println("")
	resp, err := alchemist.ImageExtraction("url", url.Values{}, demoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Image ##")
	fmt.Println("Image:", resp["image"])
}

func imageTaggingExample(alchemist *alchemyAPI.Alchemist, imageURL string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#          Image Tagging Example           #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing url: ", imageURL)
	fmt.Println("")
	resp, err := alchemist.ImageTagging("url", url.Values{}, imageURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Tags ##")
	for _, tag := range resp["imageKeywords"].([]interface{}){
		fmt.Println("keyword:", tag.(map[string] interface{})["text"])
		fmt.Println("score:", tag.(map[string] interface{})["score"])
	}
}

func taxonomyExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#            Taxonomy Example              #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Taxonomy("text", url.Values{}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)
	fmt.Println("")
	fmt.Println("## Categories ##")
	for _, category := range resp["taxonomy"].([]interface{}){
		fmt.Println("label:", category.(map[string] interface{})["label"])
		fmt.Println("score:", category.(map[string] interface{})["score"])
	}
}

func combinedExample(alchemist *alchemyAPI.Alchemist, demoText string) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("############################################")
	fmt.Println("#            Combined Example              #")
	fmt.Println("############################################")
	fmt.Println("")  
	fmt.Println("")

	fmt.Println("Processing text: ", demoText)
	fmt.Println("")
	resp, err := alchemist.Combined("text", url.Values{}, demoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	Json, _ := json.Marshal(resp)
	var out bytes.Buffer
	json.Indent(&out, Json, "=", "\t")
	fmt.Println("## Response ##")
	out.WriteTo(os.Stdout)

	fmt.Println("")
	fmt.Println("## Entitites ##")
	for _, keyword := range resp["entities"].([]interface{}){
		fmt.Println("text:", keyword.(map[string] interface{})["text"])
	 	fmt.Println("type:", keyword.(map[string] interface{})["type"])
		fmt.Println("relevance:", keyword.(map[string] interface{})["relevance"])
	}

	fmt.Println("")
	fmt.Println("## Keywords ##")
	for _, keyword := range resp["keywords"].([]interface{}){
		fmt.Println("text:", keyword.(map[string] interface{})["text"])
		fmt.Println("relevance:", keyword.(map[string] interface{})["relevance"])
	}

	fmt.Println("")
	fmt.Println("## Concepts ##")
	for _, keyword := range resp["concepts"].([]interface{}){
		fmt.Println("text:", keyword.(map[string] interface{})["text"])
		fmt.Println("relevance:", keyword.(map[string] interface{})["relevance"])
	}
}


func printAscii() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("            ,                                                                                                                              ") 
	fmt.Println("      .I7777~                                                                                                                              ")
	fmt.Println("     .I7777777                                                                                                                             ")
	fmt.Println("   +.  77777777                                                                                                                            ")
	fmt.Println(" =???,  I7777777=                                                                                                                          ")
	fmt.Println("=??????   7777777?   ,:::===?                                                                                                              ")
	fmt.Println("=???????.  777777777777777777~         .77:    ??           :7                                              =$,     :$$$$$$+  =$?          ")
	fmt.Println(" ????????: .777777777777777777         II77    ??           :7                                              $$7     :$?   7$7 =$?          ")
	fmt.Println("  .???????=  +7777777777777777        .7 =7:   ??   :7777+  :7:I777?    ?777I=  77~777? ,777I I7      77   +$?$:    :$?    $$ =$?          ")
	fmt.Println("    ???????+  ~777???+===:::         :7+  ~7   ?? .77    +7 :7?.   II  7~   ,I7 77+   I77   ~7 ?7    =7:  .$, =$    :$?  ,$$? =$?          ")
	fmt.Println("    ,???????~                        77    7:  ?? ?I.     7 :7     :7 ~7      7 77    =7:    7  7    7~   7$   $=   :$$$$$$~  =$?          ")
	fmt.Println("    .???????  ,???I77777777777~     :77777777~ ?? 7:        :7     :7 777777777:77    =7     7  +7  ~7   $$$$$$$$I  :$?       =$?          ")
	fmt.Println("   .???????  ,7777777777777777      7=      77 ?? I+      7 :7     :7 ??      7,77    =7     7   7~ 7,  =$7     $$, :$?       =$?          ")
	fmt.Println("  .???????. I77777777777777777     +7       ,7???  77    I7 :7     :7  7~   .?7 77    =7     7   ,77I   $+       7$ :$?       =$?          ")
	fmt.Println(" ,???????= :77777777777777777~     7=        ~7??  ~I77777  :7     :7  ,777777. 77    =7     7    77,  +$        .$::$?       =$?          ")
	fmt.Println(",???????  :7777777                                                                                77                                       ")
	fmt.Println(" =?????  ,7777777                                                                               77=                                        ")
	fmt.Println("   +?+  7777777?                                                                                                                           ")
	fmt.Println("    +  ~7777777                                                                                                                            ")
	fmt.Println("       I777777                                                                                                                             ")
	fmt.Println("          :~                                                                                                                               ")
}

func main() {
	keyBytes, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	alchemist := alchemyAPI.NewAlchemist(strings.NewReader(string(keyBytes[:40])))
	demoText := "Yesterday dumb Bob destroyed my fancy iPhone in beautiful Denver, Colorado. I guess I will have to head over to the Apple Store and buy a new one."
	demoHTML := "<html><head><title>Golang Demo | AlchemyAPI</title></head><body><h1>Did you know that AlchemyAPI works on HTML?</h1><p>Well, you do now.</p></body></html>"
	demoURL := "http://www.npr.org/2013/11/26/247336038/dont-stuff-the-turkey-and-other-tips-from-americas-test-kitchen"
	imageURL := "http://demo1.alchemyapi.com/images/vision/football.jpg"
	printAscii()
	entitiesExample(alchemist, demoText)
	keywordsExample(alchemist, demoText)
	conceptsExample(alchemist, demoText)
	sentimentExample(alchemist, demoHTML)
	targetedSentimentExample(alchemist, demoText)
	textExtractionExample(alchemist, demoURL)
	authorExtractionExample(alchemist, demoURL)
	languageDetectionExample(alchemist, demoText)
	titleExtractionExample(alchemist, demoURL)
	relationsExample(alchemist, demoText)
	textCategorizationExample(alchemist, demoText)
	feedDetectionExample(alchemist, demoURL)
	microformatsExample(alchemist, demoURL)
	imageExtractionExample(alchemist, demoURL)
	imageTaggingExample(alchemist, imageURL)
	taxonomyExample(alchemist, demoText)
	combinedExample(alchemist, demoText)
}
