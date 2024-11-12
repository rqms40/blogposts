package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readLine(titlePrefix)
	description := readLine(descriptionPrefix)
	tags := strings.Split(readLine(tagsPrefix), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}
}
