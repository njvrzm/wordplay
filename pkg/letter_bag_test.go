package wordplay

import "testing"

func TestLetterBagIsSubset(t *testing.T) {
	yes_cases := []struct {
		sub, sup string
	}{
		{"hell", "hello"},
		{"yes", "heresy"},
		{"stamp", "antipostmodern"},
	}
	for _, c := range yes_cases {
		if !LetterBagFromString(c.sub).IsSubset(LetterBagFromString(c.sup)) {
			t.Errorf("%q should be a subset of %q", c.sub, c.sup)
		}
	}
}
func TestLetterBagIsNotSubset(t *testing.T) {
	no_cases := []struct {
		sub, sup string
	}{
		{"diamond", "hello"},
		{"aa", "heresy"},
		{"postmoderity", "antipostmodern"},
	}
	for _, c := range no_cases {
		if LetterBagFromString(c.sub).IsSubset(LetterBagFromString(c.sup)) {
			t.Errorf("%q should not be a subset of %q", c.sub, c.sup)
		}
	}
}

func TestWordMinus(t *testing.T) {
	cases := []struct {
		word, subtrahend, desired string
	}{
		{"word", "do", "wr"},
		{"stipulation", "pistol", "uatin"},
		{"susurrus", "russ", "urus"},
	}
	for _, c := range cases {
		three := LetterBagFromString(c.word)
		two := LetterBagFromString(c.subtrahend)
		one := LetterBagFromString(c.desired)

		if three.Minus(two) != one {
			t.Errorf("%q - %q should be %q but got %q", c.word, c.subtrahend, one, three.Minus(two))
		}
	}
}
