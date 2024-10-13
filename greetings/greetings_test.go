package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	want := regexp.MustCompile(`\bOliver\b`)
	msg, err := Hello("Oliver")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Oliver") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}