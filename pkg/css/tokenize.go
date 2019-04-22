package css

import "fmt"

const (
	// `'`
	SINGLE_QUOTE = '\''
	// "
	DOUBLE_QUOTE = '"'
	// `\`
	BACKSLASH = '\\'
	// /
	SLASH = '/'
	// `\n`
	NEWLINE = '\n'
	// ` `
	SPACE = ' '
	// `\f`
	FEED = '\f'
	// `\t`
	TAB = '\t'
	// `\r`
	CR = '\r'
	// `[`
	OPEN_SQUARE = '['
	// `]`
	CLOSE_SQUAER = ']'
	// `(`
	OPEN_PARENTHESES = '('
	// `)`
	CLOSE_PARENTHESES = ')'
	// `{`
	OPEN_CURLY = '{'
	// `}`
	CLOSE_CURLY = '}'
	// `;`
	SEMICOLON = ';'
	// `*`
	ASTERISK = '*'
	// `:`
	COLON = ':'
	// `@`
	AT = '@'
)

// Position token on text line: column
type Position struct {
	Line   uint32
	Column uint32
}

func (p *Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

type Token interface {
	Start() *Position
	End() *Position
	Source() []byte
}

type BaseToken struct {
	source []byte
	start *Position
	end *Position
}

func (b *BaseToken) Start() *Position  {
	return b.start
}


func (b *BaseToken) End() *Position  {
	return b.end
}

func (b *BaseToken) Source() []byte {
	return b.source
}

type CommentToken struct {
	BaseToken
}

type DeclarationToken struct {
	BaseToken
}
