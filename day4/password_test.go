package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	cases := []struct {
		s     string
		valid bool
	}{
		{"111111", true},
		{"111123", true},
		{"223450", false},
		{"123789", false},
		{"654321", false},
		{"554321", false},
		{"655432", false},
		{"234489", true},
	}

	for _, c := range cases {
		not := " not"
		if c.valid {
			not = ""
		}

		t.Run(fmt.Sprintf("password %s is%s valid", c.s, not), func(t *testing.T) {
			if c.valid {
				assert.True(t, IsValidPassword(c.s))
			} else {
				assert.False(t, IsValidPassword(c.s))
			}
		})
	}
}
