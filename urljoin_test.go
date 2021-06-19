package urljoin

import (
	"fmt"
	"testing"
)

type MultiTestWrapper struct {
	name      string
	baseParam string
	params    []string
	expected  string
}

func TestUrlJoin(t *testing.T) {
	tests := []MultiTestWrapper{
		{
			name:      "test simple add",
			baseParam: "https://sample.com",
			params:    []string{"hello"},
			expected:  "https://sample.com/hello",
		},
		{
			name:      "with params",
			baseParam: "https://sample.com",
			params:    []string{"world", "?value=1", "?hello=world"},
			expected:  "https://sample.com/world?value=1&hello=world",
		},
		{
			name:      "auto added http",
			baseParam: "sample.com",
			params:    []string{},
			expected:  "http://sample.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if output, err := UrlJoin(test.baseParam, test.params...); err != nil {
				t.Fatal(err)
			} else {
				if output != test.expected {
					t.Fatal(fmt.Sprintf("test: %s :: params: %s, %s    -> output: %s | expected: %s", test.name, test.baseParam, test.params, output, test.expected))
				}
			}
		})
	}

}
