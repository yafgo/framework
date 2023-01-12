package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	for _, flg := range [...]randomFlag{FLG_ALL, FLG_LETTER, FLG_NUMBER} {
		for _, v := range []int{-1, 0, 1, 2, 10, 100, 256} {
			t.Logf(`[Flg: %d] Random(%v): "%v"`, flg, v, Random(v, flg))
		}
	}

	for i := 0; i < 10; i++ {
		j := i * 2
		assert.Len(t, Random(j), j)
	}
}

func TestRandomFlag(t *testing.T) {
	for _, v := range [...]randomFlag{FLG_ALL, FLG_LETTER, FLG_NUMBER, 5, -1} {
		fmt.Println(v.String())
	}
}
