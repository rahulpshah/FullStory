package main

import "testing"

func TestMain(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{
			a:        1,
			b:        1,
			expected: 1,
		},
		{
			a:        2,
			b:        1,
			expected: 2,
		},
		{
			a:        3,
			b:        2,
			expected: 6,
		},
		{
			a:        0,
			b:        2,
			expected: 0,
		},
		{
			a:        2,
			b:        0,
			expected: 0,
		},
		{
			a:        100000,
			b:        100000,
			expected: 10000000000,
		},
	}
	for _, c := range cases {
		actual := multiply(c.a, c.b)

		if actual != c.expected {
			t.Errorf("Expected %d Got %d", c.expected, actual)
		}

	}
}
