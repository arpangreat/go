package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unicode"
)

// A token is a chunk of text from a statement with a type
type token struct {
	text string
	typ  tokenTyp
}

// A tokenTyp identifies what kind of token something is
type tokenTyp int

const (
	// A bare word is a unquoted key; like 'foo' in json.foo = 1;
	typBare tokenTyp = iota

	// Numeric key; like '2' in json[2] = "foo";
	typNumericKey

	// A quoted key; like 'foo bar' in json["foo bar"] = 2;
	typQuotedKey

	// Punctuation types
	typDot    // .
	typLBrace // [
	typRBrace // ]
	typEquals // =
	typSemi   // ;

	// Value types
	typString      // "foo"
	typNumber      // 4
	typTrue        // true
	typFalse       // false
	typNull        // null
	typEmptyArray  // []
	typEmptyObject // {}

	// Ignored token
	typIgnored

	// Error token
	typError
)

// isValue returns true if the token is a valid value type
func (t token) isValue() bool {
	switch t.typ {
	case typString, typNumber, typTrue, typFalse, typNull, typEmptyArray, typEmptyObject:
		return true
	default:
		return false
	}
}

// isPunct returns true is the token is a punctuation type
func (t token) isPunct() bool {
	switch t.typ {
	case typDot, typLBrace, typRBrace, typEquals, typSemi:
		return true
	default:
		return false
	}
}

// format returns the formatted version of the token text
func (t token) format() string {
	if t.typ == typEquals {
		return " " + t.text + " "
	}
	return t.text
}

// valueTokenFromInterface takes any valid value and
// returns a value token to represent it
func valueTokenFromInterface(v interface{}) token {
	switch vv := v.(type) {

	case map[string]interface{}:
		return token{"{}", typEmptyObject}
	case []interface{}:
		return token{"[]", typEmptyArray}
	case json.Number:
		return token{vv.String(), typNumber}
	case string:
		return token{quoteString(vv), typString}
	case bool:
		if vv {
			return token{"true", typTrue}
		}
		return token{"false", typFalse}
	case nil:
		return token{"null", typNull}
	default:
		return token{"", typError}
	}
}

// quoteString takes a string and returns a quoted and
// escaped string valid for use in gron output
func quoteString(s string) string {

	out := &bytes.Buffer{}
	// bytes.Buffer never returns errors on these methods.
	// errors are explicitly ignored to keep the linter
	// happy. A price worth paying so that the linter
	// remains useful.
	_ = out.WriteByte('"')

	for _, r := range s {

		if r == '\\' || r == '"' {
			_ = out.WriteByte('\\')
			_, _ = out.WriteRune(r)
			continue
		}

		// \u2028 and \u2029 are separator runes that are not valid
		// in javascript strings so they must be escaped.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset
		if r == '\u2028' {
			_, _ = out.WriteString(`\u2028`)
			continue
		}
		if r == '\u2029' {
			_, _ = out.WriteString(`\u2029`)
			continue
		}

		// Any other control runes must be escaped
		if unicode.IsControl(r) {

			switch r {
			case '\b':
				_ = out.WriteByte('\\')
				_ = out.WriteByte('b')
			case '\f':
				_ = out.WriteByte('\\')
				_ = out.WriteByte('f')
			case '\n':
				_ = out.WriteByte('\\')
				_ = out.WriteByte('n')
			case '\r':
				_ = out.WriteByte('\\')
				_ = out.WriteByte('r')
			case '\t':
				_ = out.WriteByte('\\')
				_ = out.WriteByte('t')
			default:
				_, _ = out.WriteString(fmt.Sprintf(`\u%04X`, r))
			}

			continue
		}

		// Unescaped rune
		_, _ = out.WriteRune(r)
	}

	_ = out.WriteByte('"')
	return out.String()

}
