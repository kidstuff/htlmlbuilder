// Package htmlbuilder provides a simple way to build HTML documents.
// Here we have all HTML 5 tags.
package htlmlbuilder

func HTML(children ...*Element) *Element {
	return NewElement("html", children...)
}

func Head(children ...*Element) *Element {
	return NewElement("head", children...)
}
