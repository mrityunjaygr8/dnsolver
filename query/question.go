package query

import (
	"bytes"
	"log"
)

type DNSQuestion struct {
	Name  []byte
	Type  uint16
	Class uint16
}

func (q DNSQuestion) ToBytes() []byte {
	bs := new(bytes.Buffer)
	err := writeBytesBE(bs, q.Name)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}

	err = writeBytesBE(bs, q.Type)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}
	err = writeBytesBE(bs, q.Class)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}
	return bs.Bytes()
}
