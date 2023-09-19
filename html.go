package htmlbuilder

import (
	"fmt"
	"strings"
)

type Html struct {
	builder *builder
	name    string
	level   int
}

func getBuilder(h *Html) *builder {
	if h.builder == nil {
		panic(fmt.Sprintf("element '%s' is closed, do not call anything on it", h.name))
	}
	return h.builder
}

func (h *Html) Text(str string) *Html {
	getBuilder(h).text(str)
	return h
}

func NewHtml() *Html {
	b := &builder{Builder: &strings.Builder{}}
	h := &Html{builder: b}

	// Cyclic dependency :-(
	b.html = h
	return h
}

// create new element
func (h *Html) El(name string, attrs ...string) *Html {
	return getBuilder(h).newElement(h.level, name, attrs)
}

func (h *Html) ElSingle(name string, attrs ...string) *Html {
	getBuilder(h).newElementSingle(h.level, name, attrs)
	return h
}

func (h *Html) Attr(attrs ...string) *Html {
	getBuilder(h).attrs(attrs)
	return h
}

func (h *Html) C(body func(hb *Html)) *Html {
	body(h)
	return h.Close()
}

func (h *Html) Close() *Html {
	l := h.level
	if l > 0 {
		l = l - 1
	}
	parent := getBuilder(h).close(l)
	if parent != nil {
		return parent
	}
	return h
}

func (h *Html) String() string {
	return getBuilder(h).String()
}
