package counter

import (
	// ↓これすごい。大助かり！
	"github.com/yuin/charsetutil"
)

// Count はsをcharsetに変換してByte数を返却する
func Count(s string, charset string) (int, error) {
	b, err := charsetutil.EncodeString(s, charset)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}
