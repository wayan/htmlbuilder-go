package htmlbuilder

import (
	"html"
	"strings"
)

type builder struct {
	*strings.Builder
	html        *Html
	elems       []*Html
	openForAttr bool
}

func (b *builder) closeAttr() {
	if b.openForAttr {
		b.write(">")
		b.openForAttr = false
	}
}

func (b *builder) text(str string) {
	b.closeAttr()
	b.write(html.EscapeString(str))
}

func (b *builder) attrs(attrs []string) {
	if !b.openForAttr {
		panic("No element is open, cannot add attributes")
	}
	l := len(attrs)
	for i := 0; i < l; i = i + 2 {
		// if attribute is not present or = "" then value is ommited like in
		// <option selected>
		b.write(" ", attrs[i])
		if i+1 < l {
			v := attrs[i+1]
			if v != "" {
				b.write(`="`, html.EscapeString(v), `"`)
			}
		}
	}
}

func (b *builder) newElement(level int, name string, attrs []string) *Html {
	b.close(level)

	h := &Html{level: level + 1, builder: b, name: name}
	b.elems = append(b.elems, h)
	b.write("<", name)
	b.openForAttr = true
	b.attrs(attrs)

	return h
}

func (b *builder) newElementSingle(level int, name string, attrs []string) {
	b.close(level)

	b.write("<", name)
	b.openForAttr = true
	b.attrs(attrs)
}

func (b *builder) write(strs ...string) {
	for _, str := range strs {
		b.WriteString(str)
	}
}

func (b *builder) close(level int) *Html {
	b.closeAttr()

	// closing elements
	l := len(b.elems)
	for i := l - 1; i >= level; i = i - 1 {
		el := b.elems[i]
		b.write("</", el.name, ">")
		el.builder = nil
	}
	b.elems = b.elems[0:level]
	if level == 0 {
		return b.html
	}
	return b.elems[level-1]
}
