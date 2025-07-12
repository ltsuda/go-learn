package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "hello://some-site.com"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://blog.gypsydave5.com/",
		"hello://some-site.com",
	}
	expected := map[string]bool{
		"https://google.com":           true,
		"https://blog.gypsydave5.com/": true,
		"hello://some-site.com":        false,
	}

	result := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
