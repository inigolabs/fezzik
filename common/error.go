package common

import (
	"fmt"
	"os"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Check(err error) {
	_, debug := os.LookupEnv("DEBUG")

	if err != nil {
		switch e := err.(type) {
		case *gqlerror.Error:
			if e != nil {
				PrefixPretty("error", e)
				if debug {
					panic(e)
				} else {
					os.Exit(1)
				}
			}
		case gqlerror.List:
			if e != nil {
				PrefixPretty("error", e)
				if debug {
					panic(e)
				} else {
					os.Exit(1)
				}
			}
		default:
			if debug {
				panic(err)
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
