package ast

import "bytes"

type Range interface {
	Start() int
	End() int
}

type Node interface {
	// Range
	TokenLiteral() string
	String() string
	Line() int
}

//divide this into When Statements, Give Statements, Then Statements, etc
type Program struct {
	Statements []Node
}

func (program Program) String() string {
	var out bytes.Buffer
	for _, s := range program.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type WhenStatement struct {
	line int
	statements []Node
}

func (wh *WhenStatement) Line() int {
	return wh.line
}
func (wh *WhenStatement) TokenLiteral() string {
	return "When"
}
func (wh *WhenStatement) String() string {
	var out bytes.Buffer
	for _, s := range wh.statements {
		out.WriteString(s.String())
	}

	return out.String()
}
