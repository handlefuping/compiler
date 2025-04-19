package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	token1 := Parser("<div>1</div>")
	token1Valid := []Token{
		{
			Type: "tag",
			Name: "div",
		},
		{
			Type:    "text",
			Content: "1",
		},
		{
			Type: "tagEnd",
			Name: "div",
		},
	}

	assert.Equal(t, token1, token1Valid, "they should be equal")

	token2 := Parser("<div>1<span>2</span></div>")
	token2Valid := []Token{
		{
			Type: "tag",
			Name: "div",
		},
		{
			Type:    "text",
			Content: "1",
		},
		{
			Type: "tag",
			Name: "span",
		},
		{
			Type:    "text",
			Content: "2",
		},
		{
			Type: "tagEnd",
			Name: "span",
		},
		{
			Type: "tagEnd",
			Name: "div",
		},
	}

	assert.Equal(t, token2, token2Valid, "they should be equal")

}
