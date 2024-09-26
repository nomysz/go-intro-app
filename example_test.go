package main

import "testing"

func TestApiUrl(t *testing.T) {
	expected := "https://www.7timer.info/bin/astro.php?lon=1.2346&lat=8.7654&ac=0&unit=metric&output=json&tzshift=0"
	result := getAPIUrl(1.2345678, 8.7654321)
	if result != expected {
		t.Errorf("%s does not match %s", result, expected)
	}
}
