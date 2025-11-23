// Package pandoc provides Go types and functions for working with Pandoc's JSON AST format.
// It supports parsing Pandoc JSON output and provides structured access to document elements.
package pandoc

// KeyVal represents a key-value pair in Pandoc attributes.
type KeyVal struct {
	Key string
	Val string
}

// Attr represents Pandoc attributes with an identifier, classes, and key-value pairs.
// This is used for blocks and inline elements that support attributes.
type Attr struct {
	Identifier string     // Element identifier for cross-referencing
	Classes    []string   // CSS-style class names
	KeyVals    []*KeyVal  // Additional key-value attributes
}

// KeyValMap converts the key-value pairs to a map for convenient lookup.
func (a *Attr) KeyValMap() map[string]string {
	ret := make(map[string]string, len(a.KeyVals))
	for _, kv := range a.KeyVals {
		ret[kv.Key] = kv.Val
	}
	return ret
}

// HasClass checks if the attributes contain the specified class name.
func (a *Attr) HasClass(s string) bool {
	for _, c := range a.Classes {
		if c == s {
			return true
		}
	}
	return false
}

// Block is the interface implemented by all Pandoc block-level elements.
// Block elements include paragraphs, headers, lists, tables, and other
// structural components of a document.
type Block interface {
}

// BlockList is a slice of Block elements.
type BlockList []Block

// Plain represents a plain text block without paragraph formatting.
type Plain struct {
	Inlines InlineList
}

// Para represents a paragraph block element.
type Para struct {
	Inlines InlineList
}

// LineBlock represents a block of lines where line breaks are preserved.
type LineBlock struct {
	Lines []InlineList
}

// CodeBlock represents a code block with optional attributes and language specification.
type CodeBlock struct {
	Attr Attr   // Attributes including language class
	Text string // Code content
}

// RawBlock represents a raw block in a specific format (e.g., HTML, LaTeX).
type RawBlock struct {
	Format string
	Text   string
}

// BlockQuote represents a block quotation.
type BlockQuote struct {
	Blocks BlockList
}

// OrderedList represents a numbered list.
type OrderedList struct {
	StartNumber int
	NumberStyle string
	NumberDelim string
	Items       []BlockList
}

type BulletList struct {
	Items []BlockList
}

type DefinitionItem struct {
	Term        InlineList
	Definitions []BlockList
}

type DefinitionList struct {
	Items []*DefinitionItem
}

// Header represents a heading with level, attributes, and inline content.
type Header struct {
	Level   int        // Heading level (1-6)
	Attr    Attr       // Heading attributes
	Inlines InlineList // Heading text
}

// HorizontalRule represents a horizontal rule (thematic break).
type HorizontalRule struct {
}

// ColSpec specifies alignment and width for a table column.
type ColSpec struct {
	Alignment string
	ColWidth  float32
}

// Table represents a table with caption, column specifications, header, body, and footer.
type Table struct {
	Attr         Attr            // Table attributes
	ShortCaption InlineList      // Short caption for list of tables
	Caption      BlockList       // Full caption blocks
	ColSpecs     []*ColSpec      // Column specifications
	Head         TableHeadOrFoot // Table header
	Bodies       []*TableBody    // Table bodies
	Foot         TableHeadOrFoot // Table footer
}

// TableHeadOrFoot represents a table header or footer section.
type TableHeadOrFoot struct {
	Attr Attr
	Rows []*Row
}

type TableBody struct {
	Attr           Attr
	RowHeadColumns int
	Rows1          []*Row
	Rows2          []*Row
}

type Row struct {
	Attr  Attr
	Cells []*Cell
}

type Cell struct {
	Attr      Attr
	Alignment string
	RowSpan   int
	ColSpan   int
	Blocks    BlockList
}

type Div struct {
	Attr   Attr
	Blocks BlockList
}

type Figure struct {
	Attr         Attr
	ShortCaption InlineList
	Caption      BlockList
	Blocks       BlockList
}

// Null represents a null block element (placeholder).
type Null struct {
}

// Inline is the interface implemented by all Pandoc inline elements.
// Inline elements include text, formatting, links, images, and other
// content that appears within block elements.
type Inline interface {
}

// InlineList is a slice of Inline elements.
type InlineList []Inline

// Str represents a text string.
type Str struct {
	Text string
}

// InlineFmt represents inline formatting types (emphasis, strong, etc.).
type InlineFmt int

const (
	Emph        = InlineFmt(iota) // Emphasis (italic)
	Underline                      // Underlined text
	Strong                         // Strong emphasis (bold)
	Strikeout                      // Strikethrough text
	Superscript                    // Superscript text
	Subscript                      // Subscript text
	SmallCaps                      // Small capitals
)

// String returns the string representation of the inline format type.
func (f InlineFmt) String() string {
	switch f {
	case Emph:
		return "Emph"
	case Underline:
		return "Underline"
	case Strong:
		return "Strong"
	case Strikeout:
		return "Strikeout"
	case Superscript:
		return "Superscript"
	case Subscript:
		return "Subscript"
	case SmallCaps:
		return "SmallCaps"
	default:
		return "UnknownSpan"
	}
}

type Formatted struct {
	Fmt     InlineFmt
	Content InlineList
}

type Quoted struct {
	QuoteType string // SingleQuote or DoubleQuote
	Content   InlineList
}

type Cite struct {
	Id      string
	Prefix  InlineList
	Suffix  InlineList
	Mode    string // AuthorInText, SuppressAuthor, NormalCitation
	NoteNum int
	Hash    int
	Content InlineList
}

type Code struct {
	Attr Attr
	Text string
}

type Space struct {
}

type SoftBreak struct {
}

type LineBreak struct {
}

type Math struct {
	Type string // DisplayMath, InlineMath
	Text string
}

type RawInline struct {
	Format string // Text
	Text   string
}

type Target struct {
	URL   string
	Title string
}

type Link struct {
	Attr    Attr
	Content InlineList
	Target  Target
}

type Image struct {
	Attr    Attr
	Content InlineList
	Target  Target
}

type Note struct {
	Blocks BlockList
}

type Span struct {
	Attr    Attr
	Content InlineList
}
