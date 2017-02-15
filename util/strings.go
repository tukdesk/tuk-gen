package util

import (
	"bytes"
	"unicode"
)

func Camel2Snake(s string) string {
	buf := bytes.NewBuffer(nil)
	afterSpace := false
	for i, c := range s {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			if i > 0 && !afterSpace {
				buf.WriteRune('_')
			}

			c = unicode.ToLower(c)
		}

		buf.WriteRune(c)
		afterSpace = unicode.IsSpace(c)
	}

	return buf.String()
}
