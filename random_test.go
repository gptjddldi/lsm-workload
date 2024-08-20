package lsm_workload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name            string
		mode            string
		expectedCharset string
	}{
		{"English mode", "english", CharEnglishAlphabetNumber},
		{"Number mode", "number", CharNumber},
		{"Default mode", "default", CharBase62},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRandomString(tt.mode)
			assert.Equal(t, tt.expectedCharset, rs.Charset)
		})
	}
}

func TestRandomKey(t *testing.T) {
	tests := []struct {
		name    string
		mode    string
		checkFn func(*testing.T, string)
	}{
		{
			name: "English mode",
			mode: "english",
			checkFn: func(t *testing.T, key string) {
				assert.True(t, len(key) <= 5)
				assert.Regexp(t, "^[a-z0-9]+$", key)
			},
		},
		{
			name: "Number mode",
			mode: "number",
			checkFn: func(t *testing.T, key string) {
				assert.Regexp(t, "^[0-9]+$", key)
			},
		},
		{
			name: "Default mode",
			mode: "default",
			checkFn: func(t *testing.T, key string) {
				assert.True(t, len(key) <= 100)
				assert.Regexp(t, "^[a-zA-Z0-9]+$", key)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRandomString(tt.mode)
			key := rs.RandomKey()
			tt.checkFn(t, key)
		})
	}
}

func TestRandomValue(t *testing.T) {
	rs := NewRandomString("default")
	for i := 0; i < 1000; i++ {
		value := rs.RandomValue()

		assert.True(t, len(value) <= 1000)
		assert.Regexp(t, "^[a-zA-Z0-9]+$", value)
	}

}
