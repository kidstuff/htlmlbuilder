package htlmlbuilder

import "bytes"

type Element struct {
	name      string
	innerText string
	buff      bytes.Buffer
	Children  []*Element
}

func (e *Element) Render() string {
	e.buff.Reset()
	e.buff.WriteString("<" + e.name + ">")
	e.buff.WriteString(e.innerText)

	for _, child := range e.Children {
		e.buff.WriteString(child.Render())
	}

	e.buff.WriteString("</" + e.name + ">")

	return e.buff.String()
}

func (e *Element) InnerText(text string) *Element {
	e.innerText = text
	return e
}

func NewElement(name string, children ...*Element) *Element {
	return &Element{
		name:     name,
		Children: children,
	}
}
