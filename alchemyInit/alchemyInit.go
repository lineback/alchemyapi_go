package main

import (
	"github.com/lineback/alchemyapi_go/alchemyAPI"
	"net/url"
	"fmt"
	"io/ioutil"
)

func main() {
	alchemyAPI.Init("d5afecb608d4153b6f86e39d987cb59e0ce99b07")
	//alchemyAPI.Print()
	resp , _ := alchemyAPI.ImageTagging("url", url.Values{}, "http://demo1.alchemyapi.com/images/vision/emaxfpo.jpg")
	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(contents))
}
