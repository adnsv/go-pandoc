package pandoc

import "testing"

// A span, `[text]{#sid .cls}` in Markdown, is [Attr, [Inline]] in pandoc's JSON.
func TestSpan(t *testing.T) {
	doc, err := NewDocument([]byte(`{"pandoc-api-version":[1,23,1],"meta":{},"blocks":[
		{"t":"Para","c":[{"t":"Span","c":[["sid",["cls"],[]],[{"t":"Str","c":"text"}]]}]}]}`))
	if err != nil {
		t.Fatal(err)
	}
	blocks, err := doc.Flow()
	if err != nil {
		t.Fatal(err)
	}
	span, ok := blocks[0].(*Para).Inlines[0].(*Span)
	if !ok {
		t.Fatalf("not a span: %#v", blocks[0])
	}
	if span.Attr.Identifier != "sid" || len(span.Attr.Classes) != 1 || span.Attr.Classes[0] != "cls" {
		t.Errorf("attr %+v", span.Attr)
	}
	if s, ok := span.Content[0].(*Str); !ok || s.Text != "text" {
		t.Errorf("content %#v", span.Content)
	}
}
