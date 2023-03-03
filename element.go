package htmlbuilder

import "bytes"

type Element struct {
	name        string
	innerText   string
	isInnerText bool
	buff        bytes.Buffer
	attr        map[string]string
	Children    []*Element
}

func (e *Element) Render() string {
	e.buff.Reset()
	if e.isInnerText {
		e.buff.WriteString(e.innerText)
		return e.buff.String()
	}

	e.buff.WriteString("<" + e.name)

	if e.attr != nil {
		for key, value := range e.attr {
			e.buff.WriteString(" " + key + "=\"" + value + "\"")
		}
	}

	e.buff.WriteString(">")

	e.buff.WriteString(e.innerText)

	for _, child := range e.Children {
		e.buff.WriteString(child.Render())
	}

	e.buff.WriteString("</" + e.name + ">")

	return e.buff.String()
}

func (e *Element) Attr(key, value string) *Element {
	if e.attr == nil {
		e.attr = make(map[string]string)
	}

	e.attr[key] = value
	return e
}

func NewElement(name string, children ...*Element) *Element {
	return &Element{
		name:     name,
		Children: children,
	}
}
