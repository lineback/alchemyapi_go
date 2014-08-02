package main

import (
	"fmt"
	"os"
)

func main() {
	if 2 != len(os.Args) {
		fmt.Println("usage:  alchemyInit API_KEY")
		os.Exit(1)
	}
	apiKeyFile, err := os.Create("api_key.txt")
	defer apiKeyFile.Close()
	if err != nil {
	 	fmt.Println(err)
		os.Exit(1)
	}
	apiKeyFile.WriteString(os.Args[1])
}
