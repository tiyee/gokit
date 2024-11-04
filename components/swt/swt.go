package swt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/tiyee/gokit/cryptor"
)

var secretKey = cryptor.GenerateRandomKey(16)

type IPayload interface {
	Ver() uint8
	Encode() ([]byte, error)
	Decode([]byte) error
}
type swt struct {
	Ver             uint8
	Length          uint16
	Content         []byte
	SignatureLength uint16
	Signature       []byte
	keyBytes        []byte
}

type Opt func(s *swt)

func SecretKey(key string) Opt {
	return func(s *swt) {
		s.keyBytes = []byte(key)
	}
}
func New[T IPayload](pl T, opts ...Opt) (string, error) {
	bs, err := pl.Encode()
	if err != nil {
		return "", err
	}
	s := &swt{
		Ver:      pl.Ver(),
		Length:   uint16(len(bs)),
		Content:  bs,
		keyBytes: secretKey,
	}
	for _, opt := range opts {
		opt(s)
	}
	cntBytes := s.ContentBytes()
	hasher := hmac.New(sha256.New, s.keyBytes)
	hasher.Write(cntBytes)
	s.Signature = hasher.Sum(nil)
	s.SignatureLength = uint16(len(s.Signature))
	return s.String(), nil
}
func Parse[T IPayload](str string, pl T, opts ...Opt) error {
	bs, err := DecodeSegment(str)
	if err != nil {
		return err
	}
	s := &swt{
		Ver:      pl.Ver(),
		keyBytes: secretKey,
	}
	for _, opt := range opts {
		opt(s)
	}
	if err = s.parse(bs); err != nil {
		return err
	}
	if err := pl.Decode(s.Content); err != nil {
		return err
	}

	if pl.Ver() != s.Ver {
		return errors.New("ver error")
	}
	if !s.Valid() {
		return errors.New("valid error")
	}
	return nil
}
func (s *swt) parse(bs []byte) error {
	reader := bytes.NewReader(bs)
	var err error
	if err = binary.Read(reader, binary.BigEndian, &s.Ver); err != nil {
		return err
	}

	if err = binary.Read(reader, binary.BigEndian, &s.Length); err != nil {
		return err
	}
	fmt.Println("d payload  content length =", s.Length)
	s.Content = make([]byte, s.Length)

	if err = binary.Read(reader, binary.BigEndian, &s.Content); err != nil {
		return err
	}
	fmt.Println("d payload content", s.Content)
	if err = binary.Read(reader, binary.BigEndian, &s.SignatureLength); err != nil {
		return err
	}
	s.Signature = make([]byte, s.SignatureLength)
	if err = binary.Read(reader, binary.BigEndian, &s.Signature); err != nil {
		return err
	}
	return nil

}
func (s *swt) ContentBytes() []byte {
	writer := bytes.NewBufferString("")
	var err error
	var bs []byte
	err = binary.Write(writer, binary.BigEndian, &s.Ver)
	if err != nil {
		return bs
	}
	err = binary.Write(writer, binary.BigEndian, &s.Length)
	if err != nil {
		return bs
	}

	err = binary.Write(writer, binary.BigEndian, &s.Content)
	if err != nil {
		return bs
	}
	return writer.Bytes()
}
func (s *swt) Valid() bool {
	bs := s.ContentBytes()
	hasher := hmac.New(sha256.New, s.keyBytes)
	hasher.Write(bs)
	signature := hasher.Sum(nil)
	fmt.Println(signature, s.Signature)
	return string(signature) == string(s.Signature)
}
func (s *swt) String() string {
	bs := s.ContentBytes()
	writer := bytes.NewBuffer(bs)
	var err error
	if err = binary.Write(writer, binary.BigEndian, &s.SignatureLength); err != nil {

		return ""
	}
	if err = binary.Write(writer, binary.BigEndian, &s.Signature); err != nil {
		return ""
	}
	return EncodeSegment(writer.Bytes())
}
