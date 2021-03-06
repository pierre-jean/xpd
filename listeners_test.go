package xpd

import "testing"

func Test_consolePrinterListener_should_crash_on_Post_without_Feed(t*testing.T) {
	postWithoutFeed := Post{Subject: "dummyPost"}
	assertPanic(t, "did not crash on Post without Feed, but it should have", func() {
		ConsolePrinterListener{}.OnDuplicates(postWithoutFeed, []Post{Post{}})
	})
}

func Test_consolePrinterListener_should_crash_on_duplicate_Post_without_Feed(t*testing.T) {
	postWithFeed := Post{Subject: "dummyPost", Feed: &Feed{Id: "dummyFeed"}}
	postWithoutFeed := Post{}

	assertPanic(t, "did not crash on Post without Feed, but it should have", func() {
		ConsolePrinterListener{}.OnDuplicates(postWithFeed, []Post{postWithoutFeed})
	})
}

func Test_consolePrinterListener_happy_path(t*testing.T) {
	postWithFeed := Post{Subject: "dummyPost", Feed: &Feed{Id: "dummyFeed"}}
	ConsolePrinterListener{}.OnDuplicates(postWithFeed, []Post{postWithFeed})
}
