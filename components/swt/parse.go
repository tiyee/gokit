package swt

import "errors"

func Parse[T IPayload](signedString string, payload T) error {
	bs, err := DecodeSegment(signedString)
	if err != nil {
		return err
	}
	if len(bs) < 34 {
		return errors.New("invalid payload")
	}
	payloadBytes := bs[1 : len(bs)-31]
	if err = payload.Decode(payloadBytes); err != nil {
		return err
	}
	return nil

}
