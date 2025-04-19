package html

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var parseStr = `{
  "Type": "root",
  "Tag": "",
  "Content": "",
  "Children": [
    {
      "Type": "Element",
      "Tag": "div",
      "Content": "",
      "Children": [
        {
          "Type": "Text",
          "Tag": "",
          "Content": "1",
          "Children": null
        },
        {
          "Type": "Element",
          "Tag": "div",
          "Content": "",
          "Children": [
            {
              "Type": "Text",
              "Tag": "",
              "Content": "1",
              "Children": null
            }
          ]
        }
      ]
    }
  ]
}`

func TestParse(t *testing.T) {
	token := Tokenize(`< div id="1">1<div>1</span></ div>`)
	_, err := Parse(token)
	if assert.Error(t, err) {
		assert.Equal(t, "is invalid html", err.Error())
	}

	token1 := Tokenize(`< div id="1">1<div>1</div></ div>`)
	node, err := Parse(token1)
	assert.Empty(t, err)
	bt, err := json.MarshalIndent(node, "", "  ")
	assert.Empty(t, err)
	assert.Equal(t, string(bt), parseStr)

}
