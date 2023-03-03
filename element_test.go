package htmlbuilder

import "testing"

func TestElementRender(t *testing.T) {
	type test struct {
		name     string
		input    *Element
		expected string
	}

	tests := []test{
		{
			name:  "Empty element",
			input: NewElement("div"),

			expected: "<div></div>",
		},
		{
			name:     "Element with inner text",
			input:    NewElement("div", Text("Hello")),
			expected: "<div>Hello</div>",
		},
		{
			name:     "Element with inner text and children",
			input:    NewElement("div", Text("Hello"), NewElement("span")),
			expected: "<div>Hello<span></span></div>",
		},
		{
			name:     "Element with inner text and children",
			input:    NewElement("div", NewElement("span", Text("Hello"))),
			expected: "<div><span>Hello</span></div>",
		},
		{
			name:     "Element with inner text and children",
			input:    NewElement("div", Text("Hello"), NewElement("span", NewElement("b", Text("Hello")))),
			expected: "<div>Hello<span><b>Hello</b></span></div>",
		},
		{
			name:     "Element with attributes",
			input:    NewElement("div").Attr("id", "test").Attr("class", "test"),
			expected: "<div id=\"test\" class=\"test\"></div>",
		},
		{
			name:     "Element with attributes and inner text",
			input:    NewElement("div", Text("Hello")).Attr("id", "test").Attr("class", "test"),
			expected: "<div id=\"test\" class=\"test\">Hello</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input.Render() != test.expected {
				t.Error("Rendered element is not correct", test.input.Render())
			}
		})
	}
}
