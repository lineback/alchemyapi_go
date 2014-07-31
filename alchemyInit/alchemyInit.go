package main

import (
	"github.com/lineback/alchemyapi_go/alchemyAPI"
	"net/url"
	"fmt"
	"strings"
)

func main() {
	alchemist := alchemyAPI.NewAlchemist(strings.NewReader("d5afecb608d4153b6f86e39d987cb59e0ce99b07"))

	resp , _ := alchemist.ImageTagging("url", url.Values{}, "http://demo1.alchemyapi.com/images/vision/emaxfpo.jpg")
	
	for _, keyword := range resp["imageKeywords"].([]interface{}){
		fmt.Println("tag:", keyword.(map[string] interface{})["text"])
		fmt.Println("score:", keyword.(map[string] interface{})["score"])
	}
}
