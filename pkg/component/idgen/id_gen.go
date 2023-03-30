package idgen

import (
	"strconv"
	"strings"
	"time"
)

type IDGen struct {
	id int64
}

func Gen(biz Biz) *IDGen {
	//1+41+6+15
	nanoTime := time.Now().UnixNano()
	ts := nanoTime/int64(time.Millisecond) - ZERO*1000
	return &IDGen{
		id: int64(ts)<<21 | int64(biz)<<15 | int64(nanoTime%32768),
	}
}
func makeID(ts, biz, mask int64) int64 {
	return ts<<21 | biz<<15 | mask%32768
}
func FromId(id int64) *IDGen {
	return &IDGen{
		id: id,
	}
}
func Pow32(y int) int64 {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return 32
	}

	var ret int64 = 32
	for i := 2; i <= y; i++ {
		ret *= 32
	}
	return ret
}
func FromCode(code string) *IDGen {
	ret := Decode(code)
	return FromId(ret)
}
func (ig *IDGen) Id() int64 {
	return ig.id
}
func (ig *IDGen) DigitalStr() string {
	return strconv.FormatInt(ig.id, 10)
}
func Encode(id int64) string {
	var divisor int64 = 32
	arr := make([]string, 0)
	is := make([]int64, 0)
	for id > 0 {
		is = append(is, id%divisor)
		arr = append(arr, string(base32[id%divisor]))
		id /= divisor
	}
	return strings.Join(arr, "")
}
func (ig *IDGen) Encode() string {
	return Encode(ig.id)
}
func Decode(code string) int64 {
	m := make(map[rune]int)
	for k, v := range []rune(base32) {
		m[v] = k
	}
	arr := make([]int, 0, len(code))
	for _, v := range []rune(code) {
		if item, exist := m[v]; exist {
			arr = append(arr, item)
		}
	}
	var ret int64 = 0
	for k, v := range arr {
		ret += int64(v) * Pow32(k)
	}
	return ret
}
func (ig *IDGen) Biz() Biz {
	return Biz((ig.id >> 15) % 64)
}
func (ig *IDGen) Time() int64 {
	return ig.id>>21 + ZERO*1000
}
func (ig *IDGen) SecondsTime() int64 {
	return ig.Time() / 1000
}
