package htlmlbuilder

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
			input:    NewElement("div").InnerText("Hello"),
			expected: "<div>Hello</div>",
		},
		{
			name:     "Element with inner text and children",
			input:    NewElement("div", NewElement("span")).InnerText("Hello"),
			expected: "<div>Hello<span></span></div>",
		},
		{
			name:     "Element with inner text and children",
			input:    NewElement("div", NewElement("span").InnerText("Hello")),
			expected: "<div><span>Hello</span></div>",
		},
		{
			name:     "Element with inner text and children",
			input:    NewElement("div", NewElement("span", NewElement("b").InnerText("Hello"))).InnerText("Hello"),
			expected: "<div>Hello<span><b>Hello</b></span></div>",
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
