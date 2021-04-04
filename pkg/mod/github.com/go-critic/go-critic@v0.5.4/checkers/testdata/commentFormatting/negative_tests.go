//nolint // reason
package checker_test

//nolint reason

/*
multi-line comments
are ignored
*/

// Special kinds of comments are permitted:
//+build
//-foo

//-style comments

//directive: abc

//go:generate abc

//#ifdef foo
//#endif

//!something

//nolint

//line /foo/bar.go:10

//line /foo/bar/f-ad/a_d.go:13
//line /bar.go:14

//export myfunc
func myfunc() {
}

//go:noinline
func f2() {
	//nolint

	//	code
	//	example
	//	leading tabs

	// comment with normal style

	// this comment has empty lines
	//
	//
	// inside it.
}
