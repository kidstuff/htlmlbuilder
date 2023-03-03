// Package htmlbuilder provides a simple way to build HTML documents.
// Here we have all HTML 5 tags.
package htlmlbuilder

func HTML(children ...*Element) *Element {
	return NewElement("html", children...)
}

func Text(text string) *Element {
	return &Element{
		isInnerText: true,
		innerText:   text,
	}
}

func Abbr(children ...*Element) *Element {
	return NewElement("abbr", children...)
}

func Address(children ...*Element) *Element {
	return NewElement("address", children...)
}

func Area(children ...*Element) *Element {
	return NewElement("area", children...)
}

func Article(children ...*Element) *Element {
	return NewElement("article", children...)
}

func Aside(children ...*Element) *Element {
	return NewElement("aside", children...)
}

func Audio(children ...*Element) *Element {
	return NewElement("audio", children...)
}

func Base(children ...*Element) *Element {
	return NewElement("base", children...)
}

func Bdi(children ...*Element) *Element {
	return NewElement("bdi", children...)
}

func Bdo(children ...*Element) *Element {
	return NewElement("bdo", children...)
}

func Blockquote(children ...*Element) *Element {
	return NewElement("blockquote", children...)
}

func Body(children ...*Element) *Element {
	return NewElement("body", children...)
}

func Br(children ...*Element) *Element {
	return NewElement("br", children...)
}

func Button(children ...*Element) *Element {
	return NewElement("button", children...)
}

func Canvas(children ...*Element) *Element {
	return NewElement("canvas", children...)
}

func Caption(children ...*Element) *Element {
	return NewElement("caption", children...)
}

func Cite(children ...*Element) *Element {
	return NewElement("cite", children...)
}

func Code(children ...*Element) *Element {
	return NewElement("code", children...)
}

func Col(children ...*Element) *Element {
	return NewElement("col", children...)
}

func Colgroup(children ...*Element) *Element {
	return NewElement("colgroup", children...)
}

func Command(children ...*Element) *Element {
	return NewElement("command", children...)
}

func Datalist(children ...*Element) *Element {
	return NewElement("datalist", children...)
}

func Dd(children ...*Element) *Element {
	return NewElement("dd", children...)
}

func Del(children ...*Element) *Element {
	return NewElement("del", children...)
}

func Details(children ...*Element) *Element {
	return NewElement("details", children...)
}

func Dfn(children ...*Element) *Element {
	return NewElement("dfn", children...)
}

func Div(children ...*Element) *Element {
	return NewElement("div", children...)
}

func Dl(children ...*Element) *Element {
	return NewElement("dl", children...)
}

func Dt(children ...*Element) *Element {
	return NewElement("dt", children...)
}

func Em(children ...*Element) *Element {
	return NewElement("em", children...)
}

func Embed(children ...*Element) *Element {
	return NewElement("embed", children...)
}

func Fieldset(children ...*Element) *Element {
	return NewElement("fieldset", children...)
}

func Figcaption(children ...*Element) *Element {
	return NewElement("figcaption", children...)
}

func Figure(children ...*Element) *Element {
	return NewElement("figure", children...)
}

func Footer(children ...*Element) *Element {
	return NewElement("footer", children...)
}

func Form(children ...*Element) *Element {
	return NewElement("form", children...)
}

func H1(children ...*Element) *Element {
	return NewElement("h1", children...)
}

func H2(children ...*Element) *Element {
	return NewElement("h2", children...)
}

func H3(children ...*Element) *Element {
	return NewElement("h3", children...)
}

func H4(children ...*Element) *Element {
	return NewElement("h4", children...)
}

func H5(children ...*Element) *Element {
	return NewElement("h5", children...)
}

func H6(children ...*Element) *Element {
	return NewElement("h6", children...)
}

func Head(children ...*Element) *Element {
	return NewElement("head", children...)
}

func Header(children ...*Element) *Element {
	return NewElement("header", children...)
}

func Hr(children ...*Element) *Element {
	return NewElement("hr", children...)
}

func I(children ...*Element) *Element {
	return NewElement("i", children...)
}

func Iframe(children ...*Element) *Element {
	return NewElement("iframe", children...)
}

func Img(children ...*Element) *Element {
	return NewElement("img", children...)
}

func Input(children ...*Element) *Element {
	return NewElement("input", children...)
}

func Ins(children ...*Element) *Element {
	return NewElement("ins", children...)
}

func Kbd(children ...*Element) *Element {
	return NewElement("kbd", children...)
}

func Keygen(children ...*Element) *Element {
	return NewElement("keygen", children...)
}

func Label(children ...*Element) *Element {
	return NewElement("label", children...)
}

func Legend(children ...*Element) *Element {
	return NewElement("legend", children...)
}

func Li(children ...*Element) *Element {
	return NewElement("li", children...)
}

func Link(children ...*Element) *Element {
	return NewElement("link", children...)
}

func Main(children ...*Element) *Element {
	return NewElement("main", children...)
}

func Map(children ...*Element) *Element {
	return NewElement("map", children...)
}

func Mark(children ...*Element) *Element {
	return NewElement("mark", children...)
}

func Menu(children ...*Element) *Element {
	return NewElement("menu", children...)
}

func Menuitem(children ...*Element) *Element {
	return NewElement("menuitem", children...)
}

func Meta(children ...*Element) *Element {
	return NewElement("meta", children...)
}

func Meter(children ...*Element) *Element {
	return NewElement("meter", children...)
}

func Nav(children ...*Element) *Element {
	return NewElement("nav", children...)
}

func Noscript(children ...*Element) *Element {
	return NewElement("noscript", children...)
}

func Object(children ...*Element) *Element {
	return NewElement("object", children...)
}

func Ol(children ...*Element) *Element {
	return NewElement("ol", children...)
}

func Optgroup(children ...*Element) *Element {
	return NewElement("optgroup", children...)
}

func Option(children ...*Element) *Element {
	return NewElement("option", children...)
}

func Output(children ...*Element) *Element {
	return NewElement("output", children...)
}

func P(children ...*Element) *Element {
	return NewElement("p", children...)
}

func Param(children ...*Element) *Element {
	return NewElement("param", children...)
}

func Pre(children ...*Element) *Element {
	return NewElement("pre", children...)
}

func Progress(children ...*Element) *Element {
	return NewElement("progress", children...)
}

func Q(children ...*Element) *Element {
	return NewElement("q", children...)
}

func Rp(children ...*Element) *Element {
	return NewElement("rp", children...)
}

func Rt(children ...*Element) *Element {
	return NewElement("rt", children...)
}

func Ruby(children ...*Element) *Element {
	return NewElement("ruby", children...)
}

func S(children ...*Element) *Element {
	return NewElement("s", children...)
}

func Samp(children ...*Element) *Element {
	return NewElement("samp", children...)
}

func Script(children ...*Element) *Element {
	return NewElement("script", children...)
}

func Section(children ...*Element) *Element {
	return NewElement("section", children...)
}

func Select(children ...*Element) *Element {
	return NewElement("select", children...)
}

func Small(children ...*Element) *Element {
	return NewElement("small", children...)
}

func Source(children ...*Element) *Element {
	return NewElement("source", children...)
}

func Span(children ...*Element) *Element {
	return NewElement("span", children...)
}

func Strong(children ...*Element) *Element {
	return NewElement("strong", children...)
}

func Style(children ...*Element) *Element {
	return NewElement("style", children...)
}

func Sub(children ...*Element) *Element {
	return NewElement("sub", children...)
}

func Summary(children ...*Element) *Element {
	return NewElement("summary", children...)
}

func Sup(children ...*Element) *Element {
	return NewElement("sup", children...)
}

func Table(children ...*Element) *Element {
	return NewElement("table", children...)
}

func Tbody(children ...*Element) *Element {
	return NewElement("tbody", children...)
}

func Td(children ...*Element) *Element {
	return NewElement("td", children...)
}

func Textarea(children ...*Element) *Element {
	return NewElement("textarea", children...)
}

func Tfoot(children ...*Element) *Element {
	return NewElement("tfoot", children...)
}

func Th(children ...*Element) *Element {
	return NewElement("th", children...)
}

func Thead(children ...*Element) *Element {
	return NewElement("thead", children...)
}

func Time(children ...*Element) *Element {
	return NewElement("time", children...)
}

func Title(children ...*Element) *Element {
	return NewElement("title", children...)
}

func Tr(children ...*Element) *Element {
	return NewElement("tr", children...)
}

func Track(children ...*Element) *Element {
	return NewElement("track", children...)
}

func U(children ...*Element) *Element {
	return NewElement("u", children...)
}

func Ul(children ...*Element) *Element {
	return NewElement("ul", children...)
}

func Var(children ...*Element) *Element {
	return NewElement("var", children...)
}

func Video(children ...*Element) *Element {
	return NewElement("video", children...)
}

func Wbr(children ...*Element) *Element {
	return NewElement("wbr", children...)
}
