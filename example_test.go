package urlpath_test

import (
	"fmt"

	"github.com/ehsoc/urlpath"
)

func ExampleClean() {
	fmt.Println(urlpath.Clean("/abc/"))
	// Output: /abc
}

func ExampleDiffChild() {
	c, err := urlpath.DiffChild("/a/b", "/a/b/c")
	fmt.Printf("%s, %v", c, err)
	// Output: /c, <nil>
}
