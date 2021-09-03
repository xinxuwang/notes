package bigint

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestStr2Big(t *testing.T) {
	s := "10.3454365464534243"
	b, _ := Str2Big(s, 18)
	bs := Big2Str(b, 18)
	fmt.Println(bs)
}

func TestM(t *testing.T) {
	sb := []byte(`{"p":1.3432453543543543534542345654}`)
	type bb struct {
		P decimal.Decimal `json:"p"`
	}
	var b bb
	err := json.Unmarshal(sb, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}
