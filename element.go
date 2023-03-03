package htlmlbuilder

import "bytes"

type Element struct {
	name        string
	innerText   string
	isInnerText bool
	buff        bytes.Buffer
	Children    []*Element
}

func (e *Element) Render() string {
	e.buff.Reset()
	if e.isInnerText {
		e.buff.WriteString(e.innerText)
		return e.buff.String()
	}

	e.buff.WriteString("<" + e.name + ">")
	e.buff.WriteString(e.innerText)

	for _, child := range e.Children {
		e.buff.WriteString(child.Render())
	}

	e.buff.WriteString("</" + e.name + ">")

	return e.buff.String()
}

func NewElement(name string, children ...*Element) *Element {
	return &Element{
		name:     name,
		Children: children,
	}
}
