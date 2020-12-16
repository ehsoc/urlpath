// Package urlpath implements utility functions for manipulating url paths
package urlpath

import (
	"fmt"
	"path"
	"strings"
)

// Clean returns a 'clean' standard url path:
// 1- with leading slash and no trailing slash
// 2- blank path means root path ("/")
func Clean(urlpath string) string {
	urlpath = strings.Trim(urlpath, " ")
	if urlpath == "/" || urlpath == "" {
		return "/"
	}
	//last char cannot be a slash
	if urlpath[len(urlpath)-1] == '/' {
		urlpath = urlpath[:len(urlpath)-1]
	}
	//first char must be a slash
	if urlpath[0] != '/' {
		urlpath = "/" + urlpath
	}

	return urlpath
}

// DiffChild returns the different part of childpath from the parentpath,
// error if the childpath is not a child of parentpath
func DiffChild(parentpath, childpath string) (string, error) {
	parent := Clean(parentpath)
	child := Clean(childpath)
	if child == parent {
		//paths are equal
		return "", nil
	}
	split := strings.Split(child, parent)
	if len(split) == 2 && path.Join(parent, split[1]) == child {
		//childpath is child
		return path.Clean(split[1]), nil
	}

	return "", fmt.Errorf("urlpath/DiffChild: %q is not child of %q\n", childpath, parentpath)
}
