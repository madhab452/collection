package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name             string
		expr             string
		errAssertionFn   assert.ErrorAssertionFunc
		valueAssertionFn assert.ValueAssertionFunc
	}{
		{
			name:           "invalid sort dir",
			expr:           "foo up",
			errAssertionFn: assert.Error,
		},
		{
			name:           "invalid expression",
			expr:           "foo desc bar",
			errAssertionFn: assert.Error,
		},
		{
			name:           "success case",
			expr:           "foo desc bar asc foo.bar.baz desc",
			errAssertionFn: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Parse(tt.expr)
			tt.errAssertionFn(t, err)
			if tt.valueAssertionFn != nil {
				tt.valueAssertionFn(t, res)
			}
		})
	}
}
