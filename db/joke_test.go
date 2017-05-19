package db

import (
	"testing"
)

func TestAddJoke(t *testing.T) {
	content := "content"
	img := "imgUrl"
	id := AddJoke(content, img)
	if id <= 0 {
		t.Fatal("id <=0 ")
	}
	jokes := QueryUnUploadJokes()
	if len(jokes) == 0 {
		t.Fatal("len(jokes)==0")
	}
	find := false
	for _, joke := range jokes {
		if joke.Content == content && joke.Img == img {
			find = true
		}
	}
	if !find {
		t.Fatal("insert fail")
	}
}
