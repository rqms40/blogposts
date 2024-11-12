package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title string
}

func PostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for _, f := range dir {

		posts = append(posts, makePostFromFile(fileSystem, f))
	}
	return posts
}

func makePostFromFile(fileSystem fs.FS, f fs.DirEntry) Post {
	blogFile, _ := fileSystem.Open(f.Name())
	fileContents, _ := io.ReadAll(blogFile)

	title := strings.TrimPrefix(string(fileContents), "Title: ")
	return Post{
		Title: title,
	}
}
