package htmlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type opt struct {
	value    string
	label    string
	selected bool
}

type htmlOpts []opt

func (opts htmlOpts) Build(h *Html) {
	for _, opt := range opts {
		o := h.El("option", "value", opt.value)
		if opt.selected {
			o.Attr("selected")
		}
		o.Text(opt.value)
	}
}

func buildHead(h *Html) {
	h.El("Title").Text("Test")
}

func buildBody(h *Html) {
	d := h.El("div", "id", "options").El("div")
	d.El("h1").Text("Long & short")

	opts := []opt{{value: "1", label: "jedna"}, {value: "2", label: "2", selected: true}}
	d.El("select", "name", "ooo").C(htmlOpts(opts).Build)
	d.ElSingle("br")
}

func TestSimple(t *testing.T) {
	r := assert.New(t)
	h := NewHtml()

	div := h.El("div", "id", "1")
	r.Equal(`<div id="1"`, div.String(), "Element is written but not completed")
	r.Equal(div.String(), h.String(), "String is same on any element")
	div.Attr("class", "hidden")
	r.Equal(`<div id="1" class="hidden"`, div.String(), "Attribute can be added")
	div.Text("Nice")

	r.Equal(`<div id="1" class="hidden">Nice`, div.String(), "Text can be added")

	r.Panics(func() {
		div.Attr("lang", "en")
	}, "Element is closed for attributes")

	div.Text(", ").Text("really")
	r.Equal(`<div id="1" class="hidden">Nice, really`, div.String(), "Text can be added again, Text returns the element itself")

	_ = div.El("h2").Text(`"Mark & Spencer"`)

	// another element
	r.Equal(`<div id="1" class="hidden">Nice, really<h2>&#34;Mark &amp; Spencer&#34;`, h.String(), "Element can be nested, text is quoted")

	h.El("p")

	r.Equal(`<div id="1" class="hidden">Nice, really<h2>&#34;Mark &amp; Spencer&#34;</h2></div><p`, h.String(), "Adding element closes nested elements")

	r.Panics(func() { div.Text("What") }, "Nothing can be called on floated elements")

	h.ElSingle("br")
	h.Close()
	r.Equal(`<div id="1" class="hidden">Nice, really<h2>&#34;Mark &amp; Spencer&#34;</h2></div><p></p><br>`, h.String(), "Single element can be added")

	//h.El("html").El("head").C(buildHead).El("body").C(buildBody)
	//h.Close()
	//r.Equal("<p id=>ahoj", h.String())
}

func TestContent(t *testing.T) {
	r := assert.New(t)

	r.Equal(
		`<html><head><Title>Test</Title></head><body><div id="options"><div><h1>Long &amp; short</h1><select name="ooo"><option value="1">1</option><option value="2" selected>2</option></select><br></div></div></body></html>`,
		NewHtml().El("html").El("head").C(buildHead).El("body").C(buildBody).Close().String(),
		"Complex example",
	)

}
