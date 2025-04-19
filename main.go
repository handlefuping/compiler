package main

import (
	"compiler/html"
	"encoding/json"
	"fmt"
)

func main() {
	token := html.Tokenize(`< div id="1">1<div>1</div></ div>`)
	node, err := html.Parse(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	bt, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bt))
}
