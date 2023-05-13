package query

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type DNSHeader struct {
	Id             uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
}

func writeBytesBE(wr io.Writer, data any) error {
	err := binary.Write(wr, binary.BigEndian, data)
	if err != nil {
		return err
	}

	return nil
}

func (d DNSHeader) ToBytes() []byte {
	bs := new(bytes.Buffer)

	// bs = binary.BigEndian.AppendUint16(bs, d.Id)
	// bs = binary.BigEndian.AppendUint16(bs, d.Flags)
	// bs = binary.BigEndian.AppendUint16(bs, d.NumQuestions)
	// bs = binary.BigEndian.AppendUint16(bs, d.NumAnswers)
	// bs = binary.BigEndian.AppendUint16(bs, d.NumAuthorities)
	// bs = binary.BigEndian.AppendUint16(bs, d.NumAdditionals)

	err := writeBytesBE(bs, d.Id)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}

	err = writeBytesBE(bs, d.Flags)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}

	err = writeBytesBE(bs, d.NumQuestions)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}
	err = writeBytesBE(bs, d.NumAnswers)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}
	err = writeBytesBE(bs, d.NumAuthorities)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}
	err = writeBytesBE(bs, d.NumAdditionals)
	if err != nil {
		log.Fatal("An error has occurred:", err)
	}
	return bs.Bytes()
}
