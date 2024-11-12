package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rqms40/blogposts"
)

func main() {
	posts, err := blogposts.PostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(posts)
}
