package common

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Check(err error) {
	if err != nil {
		switch e := err.(type) {
		case *gqlerror.Error:
			if e != nil {
				PrefixPretty("error", e)
				panic(e)
			}
		case gqlerror.List:
			if e != nil {
				PrefixPretty("error", e)
				panic(e)
			}
		default:
			fmt.Printf("error:%T\n", err)
			panic(err)
		}
	}
}
