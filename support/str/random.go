package str

import (
	"math/rand"
	"time"
)

const (
	charsStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsNum = "0123456789"
	charsAll = charsStr + charsNum
)

type randomFlag int

func (r randomFlag) String() string {
	arr := [...]string{charsAll, charsStr, charsNum}
	if r < 0 || int(r) > len(arr)-1 {
		return charsAll
	}
	return arr[r]
}

const (
	// Letters && Numbers
	FLG_ALL randomFlag = iota
	// Only Letters
	FLG_LETTER
	// Only Numbers
	FLG_NUMBER
)

func Random(length int, flg ...randomFlag) (ret string) {
	if length < 1 {
		return
	}
	_flg := FLG_ALL
	if len(flg) > 0 {
		_flg = flg[0]
	}
	chars := _flg.String()

	var b = make([]byte, length)
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}
