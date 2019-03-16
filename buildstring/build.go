package buildstring

import (
	"strings"
)

func Build(a string, b string, c string) string {
	var strArr []string
	strArr = append(strArr, a)
	strArr = append(strArr, b)
	strArr = append(strArr, c)
	return strings.Join(strArr, " ");
}
