package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderBibTex(bttype, id string, keyvals [][2]string) string {
	var bld strings.Builder
	bld.WriteString(fmt.Sprintf("@%s{%s,\n", bttype, id))
	var maxLen int
	for _, kv := range keyvals {
		l := len(kv[0])
		if l > maxLen {
			maxLen = l
		}
	}
	f := "  %-" + strconv.Itoa(maxLen+2) + "s = \"%s\",\n"
	for _, kv := range keyvals {
		bld.WriteString(fmt.Sprintf(f, kv[0], kv[1]))
	}
	bld.WriteString("}\n")
	return bld.String()
}

func MakeAuthorString(authors []string) string {
	return strings.Join(authors, " and ")
}
