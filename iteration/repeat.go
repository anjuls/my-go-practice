package iteration

import "strings"

func Repeat(str string, iter int) string {
	// var repeated string
	// for i := 0; i < iter; i++ {
	// 	repeated = repeated + str
	// }
	// return repeated

	return strings.Repeat(str, iter)
}
