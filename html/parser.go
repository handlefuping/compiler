package html

type Status int

// <div>
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
	Attribute
)

func isLetter(s rune) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

type Token struct {
	Type    string // tag | text | tagEnd
	Name    string
	Content string
}

func Parser(str string) []Token {
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
			if s == '>' {
				result = append(result, Token{
					Type: "tag",
					Name: string(char),
				})
				char = []rune{}
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
					Name: string(char),
				})
				char = []rune{}
				status = TagClose
			}
		}
	}
	return result
}
