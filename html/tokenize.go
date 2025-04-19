package html

type Status int

// <div class="foo" id="1">
//   <h1>
//     hello word
//   </h1>
// </div>

const (
	Initial Status = iota
	TagOpen
	TagName
	TagClose
	Text
	TagEndName

	AttrOpen
	Attr
	AttrEnd
)

func isLetter(s rune) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

type Token struct {
	Type    string // tag | text | tagEnd
	Tag     string
	Content string
	Attrs   []string
}

func Tokenize(str string) []Token {
	status := Initial
	result := []Token{}
	char := []rune{}
	for _, s := range str {
		switch status {
		case Initial:
			if s == '<' {
				status = TagOpen
			} else {
				status = Text
				char = append(char, s)
			}
		case TagOpen:
			if isLetter(s) {
				status = TagName
				char = append(char, s)
			}
			if s == '/' {
				status = TagEndName
			}
		case TagName:
			if isLetter(s) {
				status = TagName
				char = append(char, s)
			}
			if s == ' ' || s == '>' {
				result = append(result, Token{
					Type: "tag",
					Tag:  string(char),
				})
				char = []rune{}
				if s == ' ' {

					status = AttrOpen
				}
				if s == '>' {
					status = TagClose
				}

			}
		case AttrOpen:
			if isLetter(s) {
				status = Attr
				char = append(char, s)
			}
			if s == '>' {
				status = TagClose
			}
		case Attr:
			if s == ' ' || s == '>' {
				currentToken := &result[len(result)-1]
				currentToken.Attrs = append(currentToken.Attrs, string(char))
				char = []rune{}
				if s == ' ' {
					status = AttrEnd
				}
				if s == '>' {
					status = TagClose
				}

			} else {
				char = append(char, s)

				status = Attr
			}
		case AttrEnd:
			if isLetter(s) {
				status = Attr
				char = append(char, s)
			}
			if s == '>' {
				status = TagClose
			}
		case TagClose:
			if s == '<' {
				status = TagOpen
			} else {
				status = Text
				char = append(char, s)
			}
		case Text:
			if s == '<' {
				result = append(result, Token{
					Type:    "text",
					Content: string(char),
				})
				char = []rune{}
				status = TagOpen
			} else {
				status = Text
				char = append(char, s)
			}
		case TagEndName:
			if isLetter(s) {
				status = TagEndName
				char = append(char, s)
			}
			if s == '>' {
				result = append(result, Token{
					Type: "tagEnd",
					Tag:  string(char),
				})
				char = []rune{}
				status = TagClose
			}
		}
	}
	return result
}
