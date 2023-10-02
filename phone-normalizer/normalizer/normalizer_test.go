package normalizer_test

import (
	"phone-normalizer/normalizer"
	"testing"
)

type testCase struct {
	input    string
	expected string
}

func TestNormalize(t *testing.T) {
	testCases := []testCase{
		{"1234567890", "1234567890"},
		{"123-456-7890", "1234567890"},
		{"123.456.7890", "1234567890"},
		{"(123) 456-7890", "1234567890"},
		{"123 456 7890", "1234567890"},
		{"123 456 7890 x1234", "12345678901234"},
		{"123 456 7890 ext1234", "12345678901234"},
		{"123 456 7890 ext 1234", "12345678901234"},
		{"123 456 7890 #1234", "12345678901234"},
		{"123 456 7890 # 1234", "12345678901234"},
		{"123.456.7890 x1234", "12345678901234"},
		{"123.456.7890 ext1234", "12345678901234"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := normalizer.Normalize(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
