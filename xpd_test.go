package xpd

import (
	"testing"
)

func Test_adding_to_repo(t*testing.T) {
	var repo PostRepository = newSimplePostRepository()
	repo.add(Post{})

	if len(repo.findRecent()) == 0 {
		t.Errorf("PostRepository should not be empty after post added")
	}
}

func Test_sameBodyDetector_findDuplicates_finds_same_body(t*testing.T) {
	body := "some text"
	differentBody := body + " blah"

	post := Post{Body: body}

	var repo PostRepository = newSimplePostRepository()
	repo.add(post)
	repo.add(Post{Body: differentBody})

	var detector Detector = sameBodyDetector{}
	if ! sliceEquals(detector.findDuplicates(post, []Post{post}), []Post{post}) {
		t.Errorf("same-body-detector should find only the match")
	}
}

func sliceEquals(s1 []Post, s2 []Post) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, _ := range (s1) {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}