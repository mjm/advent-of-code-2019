package day14

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleFromString(t *testing.T) {
	cases := []string{
		`157 ORE => 5 NZVS`,
		`165 ORE => 6 DCFZ`,
		`44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL`,
		`12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ`,
		`179 ORE => 7 PSHF`,
		`177 ORE => 5 HKGWZ`,
		`7 DCFZ, 7 PSHF => 2 XJWVT`,
		`165 ORE => 2 GPVTF`,
		`3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`,
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("rule from string case %d", i), func(t *testing.T) {
			r, err := RuleFromString(c)
			assert.NoError(t, err)
			assert.Equal(t, c, r.String())
		})
	}
}
