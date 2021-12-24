package common

import (
	"fmt"
	"strings"
)

type builder struct {
	b *strings.Builder
}

func NewStringBuilder() *builder {
	return &builder{
		b: new(strings.Builder),
	}
}

func (b *builder) Sprint(formatStr string, in ...interface{}) {
	b.b.WriteString(fmt.Sprintf(formatStr, in...))
}

func (b *builder) Write(in ...interface{}) {
	for _, i := range in {
		switch i := i.(type) {
		case []byte:
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
