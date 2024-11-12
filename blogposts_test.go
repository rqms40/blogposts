package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/rqms40/blogposts"
)

func TestPostsFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fileSystem := fstest.MapFS{
			"hello-tdd.md": {Data: []byte(`Title: Hello, TDD world!
Description: lol
Tags: tdd, go`)},
			// "hello-world.md": {Data: []byte("Title: Hello, World!")},
		}

		// When
		posts, err := blogposts.PostsFromFS(fileSystem)

		// Then
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("expected %d posts, got %d", len(fileSystem), len(posts))
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
			Tags:        []string{"tdd", "go"},
		})
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})
		if err == nil {
			t.Errorf("expected an error but did not get one.")
		}
	})
}

func assertPost(t testing.TB, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

type FailingFS struct {
}

func (f FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no I always fail")
}
