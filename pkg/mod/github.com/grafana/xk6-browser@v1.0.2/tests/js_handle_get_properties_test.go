package tests

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSHandleGetProperties(t *testing.T) {
	tb := newTestBrowser(t)
	p := tb.NewPage(nil)

	handle, err := p.EvaluateHandle(tb.toGojaValue(`
	() => {
		return {
			prop1: "one",
			prop2: "two",
			prop3: "three"
		};
	}
	`))
	require.NoError(t, err, "expected no error when evaluating handle")

	props, err := handle.GetProperties()
	require.NoError(t, err, "expected no error when getting properties")

	value := props["prop1"].JSONValue().String()
	assert.Equal(t, value, "one", `expected property value of "one", got %q`, value)
}
