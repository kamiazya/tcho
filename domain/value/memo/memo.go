package memo

type TextType uint

const (
	None TextType = iota
	Plain
	Markdown
)

type Memo struct {
	Type TextType
	Text string
}
