package main

import (
	"compiler/html"
	"fmt"
)

func main() {
	token := html.Parser(`< div id="1">
		1
		<div>
			1
		</div>
	</ div>`)
	fmt.Println(token)

}
