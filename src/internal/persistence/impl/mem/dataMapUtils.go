package memdb

import (
	"fmt"
	"strings"
)

func buildKey(args ...interface{}) string {
	var b strings.Builder
	for i, arg := range args {
		if ( i > 0 ) {
			b.WriteString("|")
		}
		fmt.Fprintf(&b, "%v", arg)
	}
	return b.String()
}
