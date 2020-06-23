package goping

import "testing"

func TestGoPing(t *testing.T) {
	want := "200"
	if got := GoPing("https://google.com.au"); got != want {
		t.Errorf("GoPing() = %q, want %q", got, want)
	}
}

func TestInvalidURL(t *testing.T) {
	want := "Get : unsupported protocol scheme \"\""
	if got := GoPing(""); got != want {
		t.Errorf("GoPing() = %q, want %q", got, want)
	}
}

func NotFoundPage(t *testing.T) {
	want := "404"
	if got := GoPing("http://hopefullythisurldoesntactuallyexist.com.co.za.io"); got != want {
		t.Errorf("GoPing() = %q, want %q", got, want)
	}
}
