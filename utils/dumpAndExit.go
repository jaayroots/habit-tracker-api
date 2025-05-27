package utils

import (
	"encoding/json"
	"fmt"
)

func DumpAndExit(v ...interface{}) {
	for _, item := range v {
		b, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(string(b))
		}
	}
}
