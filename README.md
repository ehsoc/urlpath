# urlpath
path utility functions for specifically manipulating url paths

## DiffChild ##
Takes two arguments `parentpath` and `childpath`. It will returns the different part of `childpath` from the `parentpath`, or error if the `childpath` is not a child of `parentpath`.

Example:
```go
	urlpath.DiffChild("/a/b", "/a/b/c")
	// Output: "/c", nil
```

## Clean ##
Returns a 'clean' standard url path:
- with leading slash and no trailing slash
- blank path means root path ("/")
  
Example:
```go
	urlpath.Clean("/abc/")
	// Output: "/abc"
```
