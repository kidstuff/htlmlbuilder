package htmlbuilder

import (
	"bytes"
	"html"
)

type Renderer interface {
	Render() string
}

type Element struct {
	name          string
	elementBuff   bytes.Buffer
	attributeBuff bytes.Buffer
	content       []Renderer
}

func (e *Element) Render() string {
	e.elementBuff.Reset()
	e.attributeBuff.Reset()

	buff := bytes.Buffer{}
	buff.WriteString("<" + e.name)

	for _, c := range e.content {
		switch c.(type) {
		case *Element:
			e.elementBuff.WriteString(c.Render())
		case *InnerText:
			e.elementBuff.WriteString(c.Render())
		case *Attribute:
			e.attributeBuff.WriteString(" " + c.Render())
		}
	}

	if e.attributeBuff.Len() > 0 {
		buff.WriteString(e.attributeBuff.String())
	}

	buff.WriteString(">")
	buff.WriteString(e.elementBuff.String())
	buff.WriteString("</" + e.name + ">")

	return buff.String()
}

func NewElement(name string, content ...Renderer) *Element {
	return &Element{
		name:    name,
		content: content,
	}
}

type Attribute struct {
	key   string
	value string
}

func (a *Attribute) Render() string {
	return a.key + "=\"" + html.EscapeString(a.value) + "\""
}

func Attr(key, value string) *Attribute {
	return &Attribute{
		key:   key,
		value: value,
	}
}

type InnerText struct {
	text     string
	isUnsafe bool
}

func (i *InnerText) Render() string {
	if i.isUnsafe {
		return i.text
	}

	return html.EscapeString(i.text)
}
