package swt

import (
	"bytes"
	"encoding/binary"
)

type BasicPayload struct {
	UserId int64
	Status uint8
	Name   string
}

// Ver return the unique id of payload.Must be unique!!!
func (p *BasicPayload) Ver() uint8 {
	return 1
}

func (p *BasicPayload) Encode() ([]byte, error) {
	writer := bytes.NewBufferString("")
	var err error
	if err = binary.Write(writer, binary.BigEndian, &p.UserId); err != nil {
		return nil, err
	}
	if err = binary.Write(writer, binary.BigEndian, &p.Status); err != nil {
		return nil, err
	}
	bs := []byte(p.Name)
	if err = binary.Write(writer, binary.BigEndian, &bs); err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}
func (p *BasicPayload) Decode(b []byte) error {
	reader := bytes.NewReader(b)
	var err error
	if err = binary.Read(reader, binary.BigEndian, &p.UserId); err != nil {
		return err
	}
	if err = binary.Read(reader, binary.BigEndian, &p.Status); err != nil {
		return err
	}
	name := make([]byte, len(b)-8-1)
	if err = binary.Read(reader, binary.BigEndian, &name); err != nil {
		return err
	}
	p.Name = string(name)
	return nil
}
