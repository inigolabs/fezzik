package fezzik

import (
	"strings"

	"github.com/jensneuse/graphql-go-tools/pkg/ast"
)

type builder struct {
	b *strings.Builder
}

func NewBuilder() *builder {
	return &builder{
		b: new(strings.Builder),
	}
}

func (b *builder) Write(in ...interface{}) {
	for _, i := range in {
		switch i := i.(type) {
		case []byte:
			b.b.Write(i)
		case ast.ByteSlice:
			b.b.Write(i)
		case string:
			b.b.WriteString(i)
		default:
			panic("in must be either []byte or string")
		}
	}
}

func (b *builder) Writeln(in ...interface{}) {
	b.Write(in...)
	b.Write("\n")
}

func (b *builder) String() string {
	return b.b.String()
}
