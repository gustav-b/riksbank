package util_test

import (
	"testing"

	"github.com/zeeraw/riksbank/swea/util"
)

func Test_ParseBool(t *testing.T) {
	cases := []struct {
		value    string
		expected bool
	}{
		{"YES", true},
		{"Y", true},
		{"1", true},
		{"TRUE", true},
		{"NO", false},
		{"N", false},
		{"0", false},
		{"FALSE", false},
		{"SOMETHING", false},
	}
	for _, c := range cases {
		t.Run(c.value, func(t *testing.T) {
			actual := util.ParseBool(c.value)
			if actual != c.expected {
				t.Errorf("expected %q to be %v, was %v", c.value, c.expected, actual)
			}
		})
	}
}
