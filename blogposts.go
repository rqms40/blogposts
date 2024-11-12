package blogposts

import (
	"io/fs"
)

func PostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := makePostFromFile(fileSystem, f.Name())
		if err != nil {
			return nil, err // todo: more clarification
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func makePostFromFile(fileSystem fs.FS, fileName string) (Post, error) {
	blogFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer blogFile.Close()

	return newPost(blogFile)
}
