package blogposts_test

import (
	blogposts "github.com/rqms40/blogposts"
	"testing"
	"testing/fstest"
)

func TestBlogPosts(t *testing.T) {

	// Given
	fs := fstest.MapFS{
		"hello-tdd.md":   {Data: []byte("Title: Hello, TDD world!")},
		"hello-world.md": {Data: []byte("Title: Hello, World!")},
	}

	// When
	posts := blogposts.PostsFromFS(fs)

	// Then
	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d", len(fs), len(posts))
	}

	expectedFirstPost := blogposts.Post{Title: "Hello, TDD world!"}
	if posts[0] != expectedFirstPost {
		t.Errorf("got %#v, want %#v", posts[0], expectedFirstPost)
	}
}
