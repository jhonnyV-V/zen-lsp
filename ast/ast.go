package ast

import "bytes"

type Range interface {
	Start() int
	End() int
}

type Node interface {
	Range
	TokenLiteral() string
	String() string
	Line() int
}

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
