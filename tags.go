// Package htmlbuilder provides a simple way to build HTML documents.
// Here we have all HTML 5 tags.
package htmlbuilder

func HTML(children ...Renderer) *Element {
	return NewElement("html", children...)
}

func Text(text string) *InnerText {
	return &InnerText{
		text:     text,
		isUnsafe: false,
	}
}

func UnsafeText(text string) *InnerText {
	return &InnerText{
		text:     text,
		isUnsafe: true,
	}
}

func A(children ...Renderer) *Element {
	return NewElement("a", children...)
}

func Svg(children ...Renderer) *Element {
	return NewElement("svg", children...)
}

func Use(children ...Renderer) *Element {
	return NewElement("use", children...)
}

func Abbr(children ...Renderer) *Element {
	return NewElement("abbr", children...)
}

func Address(children ...Renderer) *Element {
	return NewElement("address", children...)
}

func Area(children ...Renderer) *Element {
	return NewElement("area", children...)
}

func Article(children ...Renderer) *Element {
	return NewElement("article", children...)
}

func Aside(children ...Renderer) *Element {
	return NewElement("aside", children...)
}

func Audio(children ...Renderer) *Element {
	return NewElement("audio", children...)
}

func Base(children ...Renderer) *Element {
	return NewElement("base", children...)
}

func Bdi(children ...Renderer) *Element {
	return NewElement("bdi", children...)
}

func Bdo(children ...Renderer) *Element {
	return NewElement("bdo", children...)
}

func Blockquote(children ...Renderer) *Element {
	return NewElement("blockquote", children...)
}

func Body(children ...Renderer) *Element {
	return NewElement("body", children...)
}

func Br(children ...Renderer) *Element {
	return NewElement("br", children...)
}

func Button(children ...Renderer) *Element {
	return NewElement("button", children...)
}

func Canvas(children ...Renderer) *Element {
	return NewElement("canvas", children...)
}

func Caption(children ...Renderer) *Element {
	return NewElement("caption", children...)
}

func Cite(children ...Renderer) *Element {
	return NewElement("cite", children...)
}

func Code(children ...Renderer) *Element {
	return NewElement("code", children...)
}

func Col(children ...Renderer) *Element {
	return NewElement("col", children...)
}

func Colgroup(children ...Renderer) *Element {
	return NewElement("colgroup", children...)
}

func Command(children ...Renderer) *Element {
	return NewElement("command", children...)
}

func Datalist(children ...Renderer) *Element {
	return NewElement("datalist", children...)
}

func Dd(children ...Renderer) *Element {
	return NewElement("dd", children...)
}

func Del(children ...Renderer) *Element {
	return NewElement("del", children...)
}

func Details(children ...Renderer) *Element {
	return NewElement("details", children...)
}

func Dfn(children ...Renderer) *Element {
	return NewElement("dfn", children...)
}

func Div(children ...Renderer) *Element {
	return NewElement("div", children...)
}

func Dl(children ...Renderer) *Element {
	return NewElement("dl", children...)
}

func Dt(children ...Renderer) *Element {
	return NewElement("dt", children...)
}

func Em(children ...Renderer) *Element {
	return NewElement("em", children...)
}

func Embed(children ...Renderer) *Element {
	return NewElement("embed", children...)
}

func Fieldset(children ...Renderer) *Element {
	return NewElement("fieldset", children...)
}

func Figcaption(children ...Renderer) *Element {
	return NewElement("figcaption", children...)
}

func Figure(children ...Renderer) *Element {
	return NewElement("figure", children...)
}

func Footer(children ...Renderer) *Element {
	return NewElement("footer", children...)
}

func Form(children ...Renderer) *Element {
	return NewElement("form", children...)
}

func H1(children ...Renderer) *Element {
	return NewElement("h1", children...)
}

func H2(children ...Renderer) *Element {
	return NewElement("h2", children...)
}

func H3(children ...Renderer) *Element {
	return NewElement("h3", children...)
}

func H4(children ...Renderer) *Element {
	return NewElement("h4", children...)
}

func H5(children ...Renderer) *Element {
	return NewElement("h5", children...)
}

func H6(children ...Renderer) *Element {
	return NewElement("h6", children...)
}

func Head(children ...Renderer) *Element {
	return NewElement("head", children...)
}

func Header(children ...Renderer) *Element {
	return NewElement("header", children...)
}

func Hr(children ...Renderer) *Element {
	return NewElement("hr", children...)
}

func I(children ...Renderer) *Element {
	return NewElement("i", children...)
}

func Iframe(children ...Renderer) *Element {
	return NewElement("iframe", children...)
}

func Img(children ...Renderer) *Element {
	return NewElement("img", children...)
}

func Input(children ...Renderer) *Element {
	return NewElement("input", children...)
}

func Ins(children ...Renderer) *Element {
	return NewElement("ins", children...)
}

func Kbd(children ...Renderer) *Element {
	return NewElement("kbd", children...)
}

func Keygen(children ...Renderer) *Element {
	return NewElement("keygen", children...)
}

func Label(children ...Renderer) *Element {
	return NewElement("label", children...)
}

func Legend(children ...Renderer) *Element {
	return NewElement("legend", children...)
}

func Li(children ...Renderer) *Element {
	return NewElement("li", children...)
}

func Link(children ...Renderer) *Element {
	return NewElement("link", children...)
}

func Main(children ...Renderer) *Element {
	return NewElement("main", children...)
}

func Map(children ...Renderer) *Element {
	return NewElement("map", children...)
}

func Mark(children ...Renderer) *Element {
	return NewElement("mark", children...)
}

func Menu(children ...Renderer) *Element {
	return NewElement("menu", children...)
}

func Menuitem(children ...Renderer) *Element {
	return NewElement("menuitem", children...)
}

func Meta(children ...Renderer) *Element {
	return NewElement("meta", children...)
}

func Meter(children ...Renderer) *Element {
	return NewElement("meter", children...)
}

func Nav(children ...Renderer) *Element {
	return NewElement("nav", children...)
}

func Noscript(children ...Renderer) *Element {
	return NewElement("noscript", children...)
}

func Object(children ...Renderer) *Element {
	return NewElement("object", children...)
}

func Ol(children ...Renderer) *Element {
	return NewElement("ol", children...)
}

func Optgroup(children ...Renderer) *Element {
	return NewElement("optgroup", children...)
}

func Option(children ...Renderer) *Element {
	return NewElement("option", children...)
}

func Output(children ...Renderer) *Element {
	return NewElement("output", children...)
}

func P(children ...Renderer) *Element {
	return NewElement("p", children...)
}

func Param(children ...Renderer) *Element {
	return NewElement("param", children...)
}

func Pre(children ...Renderer) *Element {
	return NewElement("pre", children...)
}

func Progress(children ...Renderer) *Element {
	return NewElement("progress", children...)
}

func Q(children ...Renderer) *Element {
	return NewElement("q", children...)
}

func Rp(children ...Renderer) *Element {
	return NewElement("rp", children...)
}

func Rt(children ...Renderer) *Element {
	return NewElement("rt", children...)
}

func Ruby(children ...Renderer) *Element {
	return NewElement("ruby", children...)
}

func S(children ...Renderer) *Element {
	return NewElement("s", children...)
}

func Samp(children ...Renderer) *Element {
	return NewElement("samp", children...)
}

func Script(children ...Renderer) *Element {
	return NewElement("script", children...)
}

func Section(children ...Renderer) *Element {
	return NewElement("section", children...)
}

func Select(children ...Renderer) *Element {
	return NewElement("select", children...)
}

func Small(children ...Renderer) *Element {
	return NewElement("small", children...)
}

func Source(children ...Renderer) *Element {
	return NewElement("source", children...)
}

func Span(children ...Renderer) *Element {
	return NewElement("span", children...)
}

func Strong(children ...Renderer) *Element {
	return NewElement("strong", children...)
}

func Style(children ...Renderer) *Element {
	return NewElement("style", children...)
}

func Sub(children ...Renderer) *Element {
	return NewElement("sub", children...)
}

func Summary(children ...Renderer) *Element {
	return NewElement("summary", children...)
}

func Sup(children ...Renderer) *Element {
	return NewElement("sup", children...)
}

func Table(children ...Renderer) *Element {
	return NewElement("table", children...)
}

func Tbody(children ...Renderer) *Element {
	return NewElement("tbody", children...)
}

func Td(children ...Renderer) *Element {
	return NewElement("td", children...)
}

func Textarea(children ...Renderer) *Element {
	return NewElement("textarea", children...)
}

func Tfoot(children ...Renderer) *Element {
	return NewElement("tfoot", children...)
}

func Th(children ...Renderer) *Element {
	return NewElement("th", children...)
}

func Thead(children ...Renderer) *Element {
	return NewElement("thead", children...)
}

func Time(children ...Renderer) *Element {
	return NewElement("time", children...)
}

func Title(children ...Renderer) *Element {
	return NewElement("title", children...)
}

func Tr(children ...Renderer) *Element {
	return NewElement("tr", children...)
}

func Track(children ...Renderer) *Element {
	return NewElement("track", children...)
}

func U(children ...Renderer) *Element {
	return NewElement("u", children...)
}

func Ul(children ...Renderer) *Element {
	return NewElement("ul", children...)
}

func Var(children ...Renderer) *Element {
	return NewElement("var", children...)
}

func Video(children ...Renderer) *Element {
	return NewElement("video", children...)
}

func Wbr(children ...Renderer) *Element {
	return NewElement("wbr", children...)
}
