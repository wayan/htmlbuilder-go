# Building HTML from Go code

## Usage

```go
package main

import (
	"fmt"

	hb "github.com/wayan/htmlbuilder-go"
)

func main() {
	h := hb.NewHtml()

	// adds element with attributes (as key, value, pairs)
	div := h.El("div", "class", "hidden", "lang", "en")

	// div is open, can add attributes (conditionally)
	div.Attr("id", "first-id").Text("Special")
	// adding nested element
	div.El("h1").Text("Just a header")

	// adding another element to top level, closes all nested elements
	h.El("p").Text("Short paragraph")

	// closing all elements
	h.Close()

	// following Println yields
	// <div class="hidden" lang="en" id="first-id">Special<h1>Just a header</h1></div><p>Short paragraph</p>
	fmt.Println(h.String())

	// another example using .C (content which closes current element)
	buildHead := func(h *hb.Html) {
		h.ElSingle("meta", "charset", "utf-8")
		h.El("title").Text("Title")
	}

	buildBody := func(h *hb.Html) {
		d := h.El("div", "id", "div-one")
		d.El("h1").Text("Head")
		// nested paragraph with selected attribute
		d.El("p", "selected").Text("Some paragraph")
	}

	// following Println yields
	// <html><head><meta charset="utf-8"><title>Title</title></head><body><div id="div-one"><h1>Head</h1><p selected>Some paragraph</p></div></body></html>
	fmt.Println(
		hb.NewHtml().El("html").El("head").C(buildHead).El("body").C(buildBody).Close().String(),
	)

}

```
