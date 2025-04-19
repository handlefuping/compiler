package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tagDiv := Token{
		Type: "tag",
		Name: "div",
	}
	tagEndDiv := Token{
		Type: "tagEnd",
		Name: "div",
	}
	text := Token{
		Type:    "text",
		Content: "1",
	}
	token1 := Parser("<div>1</div>")
	token1Valid := []Token{
		tagDiv,
		text,
		tagEndDiv,
	}

	assert.Equal(t, token1, token1Valid, "they should be equal")

	token2 := Parser("<div>1<div>1</div></div>")
	token2Valid := []Token{
		tagDiv,
		text,
		tagDiv,
		text,
		tagEndDiv,
		tagEndDiv,
	}

	assert.Equal(t, token2, token2Valid, "they should be equal")

	token3 := Parser(`<div id="ss">1</div>`)
	tagDiv.Attrs = []string{`id="ss"`}
	token3Valid := []Token{
		tagDiv,
		text,
		tagEndDiv,
	}

	assert.Equal(t, token3, token3Valid, "they should be equal")

}
