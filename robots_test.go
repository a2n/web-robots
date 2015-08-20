package robots

import (
	"testing"
)

func TestIsAllowURLString(t *testing.T) {
	r := NewRobots()

	url := "https://www.google.com/search/a"
	if r.IsAllowURLString("*", url) {
		t.Error(url + " is not a allow url.")
	}

	url = "https://www.google.com/a"
	if !r.IsAllowURLString("*", url) {
		t.Error(url + " is a allow url.")
	}

	url = "https://www.yahoo.com/p/a"
	if r.IsAllowURLString("*", url) {
		t.Error(url + " is not a allow url.")
	}

	url = "https://www.yahoo.com/a"
	if !r.IsAllowURLString("*", url) {
		t.Error(url + " is a allow url.")
	}
}
