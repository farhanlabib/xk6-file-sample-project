package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testLogFormatter struct{}

func (f *testLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

func TestConsoleLogFormatter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		objects  []any
		expected string
	}{
		{objects: nil, expected: ""},
		{
			objects: []any{
				map[string]any{"one": 1, "two": "two"},
				map[string]any{"nested": map[string]any{
					"sub": float64(7.777),
				}},
			},
			expected: `{"one":1,"two":"two"} {"nested":{"sub":7.777}}`,
		},
		{
			// The first object can't be serialized to JSON, so it will be
			// skipped in the output.
			objects: []any{
				map[string]any{"one": 1, "fail": make(chan int)},
				map[string]any{"two": 2},
			},
			expected: `{"two":2}`,
		},
		{
			// Mixed objects and primitive values
			objects: []any{
				map[string]any{"one": 1},
				"someString",
				42,
			},
			expected: `{"one":1} "someString" 42`,
		},
	}

	fmtr := &consoleLogFormatter{&testLogFormatter{}}

	for _, tc := range testCases {
		var data logrus.Fields
		if tc.objects != nil {
			data = logrus.Fields{"objects": tc.objects}
		}
		out, err := fmtr.Format(&logrus.Entry{Data: data})
		require.NoError(t, err)
		assert.Equal(t, tc.expected, string(out))
	}
}
