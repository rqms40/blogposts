package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "github.com/rqms40/blogposts"
)

func TestBlogPosts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fs := fstest.MapFS{
			"hello-tdd.md":   {Data: []byte("Title: Hello, TDD world!")},
			"hello-world.md": {Data: []byte("Title: Hello, World!")},
		}

		// When
		posts, err := blogposts.PostsFromFS(fs)

		// Then
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts, got %d", len(fs), len(posts))
		}

		expectedFirstPost := blogposts.Post{Title: "Hello, TDD world!"}
		if posts[0] != expectedFirstPost {
			t.Errorf("got %#v, want %#v", posts[0], expectedFirstPost)
		}
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})
		if err == nil {
			t.Errorf("expected an error but did not get one.")
		}
	})
}

type FailingFS struct {
}

func (f FailingFS) Open(string) (fs.File, error) {
	return nil, errors.New("oh no I always fail")
}
