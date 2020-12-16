package urlpath_test

import (
	"fmt"
	"testing"

	"github.com/ehsoc/urlpath"
)

type DiffTests struct {
	root, path, want string
	err              bool
}

var diffChildTests = []DiffTests{
	{"/abc/def", "/abc/def", "", false},
	{"/abc/def", "/abc/def/ghijklmn/opqrstuvwyz1234", "/ghijklmn/opqrstuvwyz1234", false},
	{"/abc/def", "/abc/defeeeeeee", "", true},
	{"/abc/def", "/abc/def/ghi", "/ghi", false},
	{"/abc/def", "/abc/def/ghi/jkl", "/ghi/jkl", false},
	{"/abc", "/a", "", true},
	{"/", "/", "", false},
	{"/ab", "/abc", "", true},
	{"/ab", "/abc", "", true},
	{"/api/v1/123", "/abc/api/v1/123", "", true},
	{"api/v1 /123", "/api/v1 /123", "", false},
	{"/api/v1/123", "api/v1/123", "", false},
	{"", "", "", false},
}

type CleanTests struct {
	path, want string
}

var cleanTests = []CleanTests{
	{"/", "/"},
	{"/abc/", "/abc"},
	{"abc/", "/abc"},
	{"/ ", "/"},
	{" ", "/"},
	{" ", "/"},
	{"    /abc/def/gh/  ", "/abc/def/gh"},
}

func TestClean(t *testing.T) {
	for _, tt := range cleanTests {
		t.Run(fmt.Sprintf("Clean(%q)", tt.path), func(t *testing.T) {
			got := urlpath.Clean(tt.path)
			if got != tt.want {
				t.Errorf("got: %q want: %q", got, tt.want)
			}
		})
	}
}

func TestDiffChild(t *testing.T) {
	for _, tt := range diffChildTests {
		t.Run(fmt.Sprintf("DiffPath(%q,%q)", tt.root, tt.path), func(t *testing.T) {
			got, err := urlpath.DiffChild(tt.root, tt.path)
			if tt.err && err == nil {
				t.Fatalf("expecting error")
			}
			if !tt.err && err != nil {
				t.Fatalf("not expecting error")
			}
			if got != tt.want {
				t.Errorf("got: %q want: %q", got, tt.want)
			}
		})
	}
}
