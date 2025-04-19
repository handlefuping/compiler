package html

import (
	"errors"
)

type Node struct {
	Type     string // Element | Text
	Tag      string
	Content  string
	Children []*Node
}

func token2Node(token Token) *Node {
	if token.Type == "tag" || token.Type == "tagEnd" {
		return &Node{
			Type: "Element",
			Tag:  token.Tag,
		}
	} else {
		return &Node{
			Type:    "Text",
			Content: token.Content,
		}
	}

}

func Parse(token []Token) (*Node, error) {
	root := &Node{
		Type: "root",
	}
	stack := []*Node{root}
	for _, currentToken := range token {
		parent := stack[len(stack)-1]
		currentNode := token2Node(currentToken)
		if currentToken.Type == "tag" {
			parent.Children = append(parent.Children, currentNode)
			stack = append(stack, currentNode)
		}
		if currentToken.Type == "text" {
			parent.Children = append(parent.Children, currentNode)
		}
		if currentToken.Type == "tagEnd" {
			lastNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if lastNode.Tag != currentNode.Tag {
				return nil, errors.New("is invalid html")
			}
		}
	}
	return root, nil

}
