package utils

import "testing"

type input struct{
	base string
	href string
}
type testdata struct {
	name string
	input input
	expected string
}

func TestNormalizeUrl(t *testing.T) {
	tests := []testdata{
		{
			name: "absolute URL stays same",
			input: input{
				base: "https://example.com",
				href: "https://example.com/blog/path",
			},
			expected: "https://example.com/blog/path",
		},
		{
			name: "remove trailing slash",
			input: input{
				base: "https://example.com",
				href: "/blog/path/",
			},
			expected: "https://example.com/blog/path",
		},
		{
			name: "relative path",
			input: input{
				base: "https://example.com/blog/",
				href: "post",
			},
			expected: "https://example.com/blog/post",
		},
		{
			name: "root relative",
			input: input{
				base: "https://example.com/blog/",
				href: "/about",
			},
			expected: "https://example.com/about",
		},
		{
			name: "remove fragment",
			input: input{
				base: "https://example.com",
				href: "/page#section",
			},
			expected: "https://example.com/page",
		},
		{
			name: "sort query params",
			input: input{
				base: "https://example.com",
				href: "/page?b=2&a=1",
			},
			expected: "https://example.com/page?a=1&b=2",
		},
		{
			name: "remove utm params",
			input: input{
				base: "https://example.com",
				href: "/page?utm_source=google&a=1",
			},
			expected: "https://example.com/page?a=1",
		},
		{
			name: "remove fbclid",
			input: input{
				base: "https://example.com",
				href: "/page?fbclid=abc123",
			},
			expected: "https://example.com/page",
		},
		{
			name: "remove default http port",
			input: input{
				base: "http://example.com",
				href: "http://example.com:80/page",
			},
			expected: "http://example.com/page",
		},
		{
			name: "remove default https port",
			input: input{
				base: "https://example.com",
				href: "https://example.com:443/page",
			},
			expected: "https://example.com/page",
		},
		{
			name: "lowercase host and scheme",
			input: input{
				base: "HTTPS://EXAMPLE.COM",
				href: "/Page",
			},
			expected: "https://example.com/Page",
		},
		{
			name: "clean path dots",
			input: input{
				base: "https://example.com",
				href: "/a/../b/./c",
			},
			expected: "https://example.com/b/c",
		},
		{
			name: "empty href",
			input: input{
				base: "https://example.com/page",
				href: "",
			},
			expected: "https://example.com/page",
		},
	}

	for i, tc := range tests {
		actual, err := normalizeUrl(tc.input.base, tc.input.href)
		if err != nil {
			t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
		}
		if actual != tc.expected {
			t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
		}
	}
}