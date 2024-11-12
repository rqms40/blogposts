package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readLine("Title: ")
	description := readLine("Description: ")

	return Post{
		Title:       title,
		Description: description,
	}
}
