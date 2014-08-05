package alchemyAPI

import (
	"testing"
	"io/ioutil"
	"strings"
	"net/url"
	"fmt"
)

func test_Entities(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Entities ...")
	resp, respErr := alchemist.Entities("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Entities FAILED")
	} else {
		t.Log("Text Entities PASSED")
	}
	resp, respErr = alchemist.Entities("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Entities FAILED")
	} else {
		t.Log("HTML Entities PASSED")
	}
	resp, respErr = alchemist.Entities("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Entities FAILED")
	} else {
		t.Log("URL Entities PASSED")
	}
	resp, respErr = alchemist.Entities("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Entities FAILED")
	} else {
		t.Log("Bad Flavor Entities PASSED")
	}
}

func test_Keywords(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Keywords ...")
	resp, respErr := alchemist.Keywords("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Keywords FAILED")
	} else {
		t.Log("Text Keywords PASSED")
	}
	resp, respErr = alchemist.Keywords("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Keywords FAILED")
	} else {
		t.Log("HTML Keywords PASSED")
	}
	resp, respErr = alchemist.Keywords("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Keywords FAILED")
	} else {
		t.Log("URL Keywords PASSED")
	}
	resp, respErr = alchemist.Keywords("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Keywords FAILED")
	} else {
		t.Log("Bad Flavor Keywords PASSED")
	}
}

func test_Concepts(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Concepts ...")
	resp, respErr := alchemist.Concepts("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Concepts FAILED")
	} else {
		t.Log("Text Concepts PASSED")
	}
	resp, respErr = alchemist.Concepts("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Concepts FAILED")
	} else {
		t.Log("HTML Concepts PASSED")
	}
	resp, respErr = alchemist.Concepts("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Concepts FAILED")
	} else {
		t.Log("URL Concepts PASSED")
	}
	resp, respErr = alchemist.Concepts("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Concepts FAILED")
	} else {
		t.Log("Bad Flavor Concepts PASSED")
	}
}

func test_Sentiment(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Sentiment ...")
	resp, respErr := alchemist.Sentiment("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Sentiment FAILED")
	} else {
		t.Log("Text Sentiment PASSED")
	}
	resp, respErr = alchemist.Sentiment("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Sentiment FAILED")
	} else {
		t.Log("HTML Sentiment PASSED")
	}
	resp, respErr = alchemist.Sentiment("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Sentiment FAILED")
	} else {
		t.Log("URL Sentiment PASSED")
	}
	resp, respErr = alchemist.Sentiment("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Sentiment FAILED")
	} else {
		t.Log("Bad Flavor Sentiment PASSED")
	}
}

func test_TargetedSentiment(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing TargetedSentiment ...")
	resp, respErr := alchemist.TargetedSentiment("text", url.Values{}, testText, "heart")
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text TargetedSentiment FAILED")
	} else {
		t.Log("Text TargetedSentiment PASSED")
	}
	resp, respErr = alchemist.TargetedSentiment("html", url.Values{}, testHTML, "language")
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML TargetedSentiment FAILED")
	} else {
		t.Log("HTML TargetedSentiment PASSED")
	}
	resp, respErr = alchemist.TargetedSentiment("url", url.Values{}, testURL, "congress")
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL TargetedSentiment FAILED")
	} else {
		t.Log("URL TargetedSentiment PASSED")
	}
	resp, respErr = alchemist.TargetedSentiment("bad_flavor", url.Values{}, testURL, "congress")
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Enities FAILED")
	} else {
		t.Log("Bad Flavor TargetedSentiment PASSED")
	}
}

func test_Text(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Text ...")
	resp, respErr := alchemist.Text("text", url.Values{}, testText)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text Text FAILED")
	} else {
		t.Log("Text Text PASSED")
	}
	resp, respErr = alchemist.Text("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Text FAILED")
	} else {
		t.Log("HTML Text PASSED")
	}
	resp, respErr = alchemist.Text("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Text FAILED")
	} else {
		t.Log("URL Text PASSED")
	}
	resp, respErr = alchemist.Text("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Text FAILED")
	} else {
		t.Log("Bad Flavor Text PASSED")
	}
}

func test_TextRaw(t *testing.T, alchemist *Alchemist, testTextRaw, testHTML, testURL string) {
	fmt.Println("Testing TextRaw ...")
	resp, respErr := alchemist.TextRaw("text", url.Values{}, testTextRaw)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text TextRaw FAILED")
	} else {
		t.Log("Text TextRaw PASSED")
	}
	resp, respErr = alchemist.TextRaw("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML TextRaw FAILED")
	} else {
		t.Log("HTML TextRaw PASSED")
	}
	resp, respErr = alchemist.TextRaw("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL TextRaw FAILED")
	} else {
		t.Log("URL TextRaw PASSED")
	}
	resp, respErr = alchemist.TextRaw("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor TextRaw FAILED")
	} else {
		t.Log("Bad Flavor TextRaw PASSED")
	}
}

func test_Author(t *testing.T, alchemist *Alchemist, testAuthor, testHTML, testURL string) {
	fmt.Println("Testing Author ...")
	resp, respErr := alchemist.Author("text", url.Values{}, testAuthor)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text Author FAILED")
	} else {
		t.Log("Text Author PASSED")
	}
	resp, respErr = alchemist.Author("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Author FAILED")
	} else {
		t.Log("HTML Author PASSED")
	}
	resp, respErr = alchemist.Author("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Author FAILED")
	} else {
		t.Log("URL Author PASSED")
	}
	resp, respErr = alchemist.Author("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Author FAILED")
	} else {
		t.Log("Bad Flavor Author PASSED")
	}
}

func test_Title(t *testing.T, alchemist *Alchemist, testTitle, testHTML, testURL string) {
	fmt.Println("Testing Title ...")
	resp, respErr := alchemist.Title("text", url.Values{}, testTitle)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text Title FAILED")
	} else {
		t.Log("Text Title PASSED")
	}
	resp, respErr = alchemist.Title("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Title FAILED")
	} else {
		t.Log("HTML Title PASSED")
	}
	resp, respErr = alchemist.Title("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Title FAILED")
	} else {
		t.Log("URL Title PASSED")
	}
	resp, respErr = alchemist.Title("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Title FAILED")
	} else {
		t.Log("Bad Flavor Title PASSED")
	}
}

func test_Language(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Language ...")
	resp, respErr := alchemist.Language("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Language FAILED")
	} else {
		t.Log("Text Language PASSED")
	}
	resp, respErr = alchemist.Language("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Language FAILED")
	} else {
		t.Log("HTML Language PASSED")
	}
	resp, respErr = alchemist.Language("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Language FAILED")
	} else {
		t.Log("URL Language PASSED")
	}
	resp, respErr = alchemist.Language("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Language FAILED")
	} else {
		t.Log("Bad Flavor Language PASSED")
	}
}

func test_Relations(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Relations ...")
	resp, respErr := alchemist.Relations("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Relations FAILED")
	} else {
		t.Log("Text Relations PASSED")
	}
	resp, respErr = alchemist.Relations("html", url.Values{}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Relations FAILED")
	} else {
		t.Log("HTML Relations PASSED")
	}
	resp, respErr = alchemist.Relations("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Relations FAILED")
	} else {
		t.Log("URL Relations PASSED")
	}
	resp, respErr = alchemist.Relations("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Relations FAILED")
	} else {
		t.Log("Bad Flavor Relations PASSED")
	}
}

func test_Category(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Category ...")
	resp, respErr := alchemist.Category("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Category FAILED")
	} else {
		t.Log("Text Category PASSED")
	}
	resp, respErr = alchemist.Category("html", url.Values{"url":{"test"}}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Category FAILED")
	} else {
		t.Log("HTML Category PASSED")
	}
	resp, respErr = alchemist.Category("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Category FAILED")
	} else {
		t.Log("URL Category PASSED")
	}
	resp, respErr = alchemist.Category("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Category FAILED")
	} else {
		t.Log("Bad Flavor Category PASSED")
	}
}

func test_Feeds(t *testing.T, alchemist *Alchemist, testFeeds, testHTML, testURL string) {
	fmt.Println("Testing Feeds ...")
	resp, respErr := alchemist.Feeds("text", url.Values{}, testFeeds)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text Feeds FAILED")
	} else {
		t.Log("Text Feeds PASSED")
	}
	resp, respErr = alchemist.Feeds("html", url.Values{"url": {"test"}}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Feeds FAILED")
	} else {
		t.Log("HTML Feeds PASSED")
	}
	resp, respErr = alchemist.Feeds("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Feeds FAILED")
	} else {
		t.Log("URL Feeds PASSED")
	}
	resp, respErr = alchemist.Feeds("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Feeds FAILED")
	} else {
		t.Log("Bad Flavor Feeds PASSED")
	}
}

func test_Microformats(t *testing.T, alchemist *Alchemist, testMicroformats, testHTML, testURL string) {
	fmt.Println("Testing Microformats ...")
	resp, respErr := alchemist.Microformats("text", url.Values{}, testMicroformats)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text Microformats FAILED")
	} else {
		t.Log("Text Microformats PASSED")
	}
	resp, respErr = alchemist.Microformats("html", url.Values{"url": {"test"}}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Microformats FAILED")
	} else {
		t.Log("HTML Microformats PASSED")
	}
	resp, respErr = alchemist.Microformats("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Microformats FAILED")
	} else {
		t.Log("URL Microformats PASSED")
	}
	resp, respErr = alchemist.Microformats("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Microformats FAILED")
	} else {
		t.Log("Bad Flavor Microformats PASSED")
	}
}

func test_Combined(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Combined ...")
	resp, respErr := alchemist.Combined("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Combined FAILED")
	} else {
		t.Log("Text Combined PASSED")
	}
	resp, respErr = alchemist.Combined("html", url.Values{}, testHTML)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("HTML Combined FAILED")
	} else {
		t.Log("HTML Combined PASSED")
	}
	resp, respErr = alchemist.Combined("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Combined FAILED")
	} else {
		t.Log("URL Combined PASSED")
	}
	resp, respErr = alchemist.Combined("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Combined FAILED")
	} else {
		t.Log("Bad Flavor Combined PASSED")
	}
}

func test_Taxonomy(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL string) {
	fmt.Println("Testing Taxonomy ...")
	resp, respErr := alchemist.Taxonomy("text", url.Values{}, testText)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Text Taxonomy FAILED")
	} else {
		t.Log("Text Taxonomy PASSED")
	}
	resp, respErr = alchemist.Taxonomy("html", url.Values{"url":{"test"}}, testHTML)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("HTML Taxonomy FAILED")
	} else {
		t.Log("HTML Taxonomy PASSED")
	}
	resp, respErr = alchemist.Taxonomy("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL Taxonomy FAILED")
	} else {
		t.Log("URL Taxonomy PASSED")
	}
	resp, respErr = alchemist.Taxonomy("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor Taxonomy FAILED")
	} else {
		t.Log("Bad Flavor Taxonomy PASSED")
	}
}

func test_ImageExtraction(t *testing.T, alchemist *Alchemist, testImageExtraction, testHTML, testURL string) {
	fmt.Println("Testing ImageExtraction ...")
	resp, respErr := alchemist.ImageExtraction("text", url.Values{}, testImageExtraction)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text ImageExtraction FAILED")
	} else {
		t.Log("Text ImageExtraction PASSED")
	}
	resp, respErr = alchemist.ImageExtraction("html", url.Values{}, testHTML)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("HTML ImageExtraction FAILED")
	} else {
		t.Log("HTML ImageExtraction PASSED")
	}
	resp, respErr = alchemist.ImageExtraction("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL ImageExtraction FAILED")
	} else {
		t.Log("URL ImageExtraction PASSED")
	}
	resp, respErr = alchemist.ImageExtraction("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor ImageExtraction FAILED")
	} else {
		t.Log("Bad Flavor ImageExtraction PASSED")
	}
}

func test_ImageTagging(t *testing.T, alchemist *Alchemist, testText, testHTML, testURL, testImage string) {
	fmt.Println("Testing ImageTagging ...")
	resp, respErr := alchemist.ImageTagging("text", url.Values{}, testText)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Text ImageTagging FAILED")
	} else {
		t.Log("Text ImageTagging PASSED")
	}
	resp, respErr = alchemist.ImageTagging("html", url.Values{}, testHTML)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("HTML ImageTagging FAILED")
	} else {
		t.Log("HTML ImageTagging PASSED")
	}
	resp, respErr = alchemist.ImageTagging("url", url.Values{}, testURL)
	if respErr != nil || resp["status"] != "OK" {
		t.Error("URL ImageTagging FAILED")
	} else {
		t.Log("URL ImageTagging PASSED")
	}
	resp, respErr = alchemist.ImageTagging("image", url.Values{}, testImage)
	fmt.Println(resp["status"])
	if respErr != nil || resp["status"] != "OK" {
		t.Error("Image ImageTagging FAILED")
	} else {
		t.Log("Image ImageTagging PASSED")
	}
	resp, respErr = alchemist.ImageTagging("bad_flavor", url.Values{}, testURL)
	if respErr == nil || resp["status"] == "OK" {
		t.Error("Bad Flavor ImageTagging FAILED")
	} else {
		t.Log("Bad Flavor ImageTagging PASSED")
	}
}

func TestAll(t *testing.T) {
	keyBytes, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		t.Error("Entities failed. Error reading api key file")
	}
	alchemist := NewAlchemist(strings.NewReader(string(keyBytes[:40])))
	testText := "Bob broke my heart, and then made up this silly sentence to test the Golang SDK"
	testHTML := "<html><head><span class=\"author\">John</span><title>The best SDK Test | AlchemyAPI</title></head><body><h1>Hello World!</h1><p>My favorite language is Golang</p></body></html>"
	testURL  :=  "http://www.nytimes.com/2013/07/13/us/politics/a-day-of-friction-notable-even-for-a-fractious-congress.html?_r=0"
	testImage := "pigeon.jpg"

	test_Entities(t, alchemist, testText, testHTML, testURL)
	test_Keywords(t, alchemist, testText, testHTML, testURL)
	test_Concepts(t, alchemist, testText, testHTML, testURL)
	test_Sentiment(t, alchemist, testText, testHTML, testURL)
	test_TargetedSentiment(t, alchemist, testText, testHTML, testURL)
	test_Text(t, alchemist, testText, testHTML, testURL)
	test_TextRaw(t, alchemist, testText, testHTML, testURL)
	test_Author(t, alchemist, testText, testHTML, testURL)
	test_Language(t, alchemist, testText, testHTML, testURL)
	test_Title(t, alchemist, testText, testHTML, testURL)
	test_Relations(t, alchemist, testText, testHTML, testURL)
	test_Concepts(t, alchemist, testText, testHTML, testURL)
	test_Category(t, alchemist, testText, testHTML, testURL)
	test_Feeds(t, alchemist, testText, testHTML, testURL)
	test_Microformats(t, alchemist, testText, testHTML, testURL)
	test_Combined(t, alchemist, testText, testHTML, testURL)
	test_Taxonomy(t, alchemist, testText, testHTML, testURL)
	test_ImageExtraction(t, alchemist, testText, testHTML, testURL)
	test_ImageTagging(t, alchemist, testText, testHTML, testURL, testImage)
}
