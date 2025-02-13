// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package util

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/88250/lute/html"
)

// InsertElem inserts a new element value at the specified index position.
// 0 <= index <= len(a)
func InsertElem[T any](ret []T, index int, value T) []T {
	if len(ret) == index { // nil or empty slice or after last element
		return append(ret, value)
	}

	ret = append(ret[:index+1], ret[index:]...) // index < len(a)
	ret[index] = value
	return ret
}

func EscapeHTML(s string) string {
	if strings.Contains(s, "&amp;") {
		return s
	}
	if strings.ContainsAny(s, "<>\"'") {
		return html.EscapeString(s)
	}
	return s
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RemoveRedundantSpace(str string) string {
	buf := bytes.Buffer{}
	lastIsChinese := false
	lastIsSpace := false
	for _, r := range str {
		if unicode.IsSpace(r) {
			if lastIsChinese || lastIsSpace {
				continue
			}
			buf.WriteRune(' ')
			lastIsChinese = false
			lastIsSpace = true
			continue
		}

		lastIsSpace = false
		buf.WriteRune(r)
		if unicode.Is(unicode.Han, r) {
			lastIsChinese = true
			continue
		} else {
			lastIsChinese = false
		}
	}
	return buf.String()
}
