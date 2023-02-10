package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Myroslav"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Myroslav")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Myroslav") = %q, %v want match for all %#q, nil`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", err`, msg, err)
	}
}
