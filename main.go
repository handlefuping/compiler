package main

import (
	"compiler/html"
	"fmt"
)

func main() {
	token := html.Parser("< div>1</ div>")
	fmt.Println(token)

}
