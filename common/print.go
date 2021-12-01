package common

import (
	"encoding/json"
	"fmt"
)

func PrefixPretty(prefix string, object interface{}) {
	objectString, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s : %s\n", prefix, objectString)
}
