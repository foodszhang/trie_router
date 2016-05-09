package TrieRouter

import "testing"

func TestTree(t *testing.T) {
	root := CreateNode()
	root.Insert("/user/go/")
	root.Insert("/article/")
	root.Insert("/article/go/123")
	root.Insert("/user/create/")
	root.Insert("/user/go/5678")
	root.Insert("/users/go/1234")
	if root.Search("/user/") != false {
		t.Error("/user/")
	}
	if root.Search("/article/go/") != false {
		t.Error("/article/go/", root.Search("/article/go/"))
	}
	if root.Search("/user/go/") == false {
		t.Error("/user/go/")
	}
	if root.Search("/user/5678/") != false {
		t.Error("/user/5678/")
	}
	if root.Search("/user/1234/") != false {
		t.Error("/user/1234/")
	}
	if root.Search("/users/go/1234") == false {
		t.Error("/users/go/1234")
	}
}
