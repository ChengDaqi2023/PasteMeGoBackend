package util

import (
    _ "github.com/PasteUs/PasteMeGoBackend/tests"
    "testing"
)

func TestValidChecker(t *testing.T) {
    var TestCases = []struct {
        input    string
        expected string
    }{
        {"01234567", "permanent"},
        {"12345678", "permanent"},
        {"asdfqwer", "temporary"},
        {"0asdf123", "temporary"},
        {"0", ""},
        {"a", ""},
        {"asdf", "temporary"},
        {"asd", "temporary"},
        {"asdfqwerasdf", ""},
        {"1000000000", ""},
        {"dafsdf?", ""},
        {"once", "temporary"},
    }
    for i, TestCase := range TestCases {
        output, _ := ValidChecker(TestCase.input)
        if output != TestCase.expected {
            t.Errorf("Test %d | input: %s, expected: %s, output: %s\n", i, TestCase.input, TestCase.expected, output)
        }
    }
}
