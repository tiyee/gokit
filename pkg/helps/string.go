package helps

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func NonceStr() string {
	ns := time.Now().UnixNano()
	bs := md5.Sum([]byte(strconv.FormatInt(ns, 10)))
	return fmt.Sprintf("%x", bs)
}
