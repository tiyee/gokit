package component

import (
	"encoding/base64"
	"encoding/json"
	"github.com/tiyee/gokit/pkg/helps"
)

type JWT struct {
	Uid     int64 `json:"u"`
	Expired int64 `json:"e"`
}

func (j *JWT) Decrypt(data []byte) error {
	var myErr error
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(buf, data)
	if err != nil {
		return err
	}
	if bs, err := helps.DecryptByDESAndCBC(buf[:n]); err == nil {
		if err := json.Unmarshal(bs, j); err == nil {
			return nil
		} else {
			myErr = err
		}
	} else {
		myErr = err
	}
	return myErr
}
func (j *JWT) Encrypt() ([]byte, error) {
	if bs, err := json.Marshal(j); err == nil {

		if data, err := helps.EncryptByDESAndCBC(bs); err == nil {
			buf := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
			base64.StdEncoding.Encode(buf, data)
			return buf, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
