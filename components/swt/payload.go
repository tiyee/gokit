package swt

type IPayload interface {
	Encode() []byte
	Decode([]byte) error
	SignKey() []byte
}
