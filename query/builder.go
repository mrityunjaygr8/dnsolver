package query

import (
	"bytes"
	"log"
)

func encodeDNSName(domainName string) []byte {
	bs := new(bytes.Buffer)

	nameParts := bytes.Split([]byte(domainName), []byte("."))

	for _, part := range nameParts {
		// bs = binary.BigEndian.AppendUint16(bs, uint16(len(part)))
		err := writeBytesBE(bs, uint8(len(part)))
		if err != nil {
			log.Fatal("An 123 error has occurred:", err)
		}
		// bs.WriteTo = append(bs, []byte(part)...)
		err = writeBytesBE(bs, part)
		if err != nil {
			log.Fatal("An error 123 has occurred:", err)
		}
	}

	err := writeBytesBE(bs, uint8(0))
	if err != nil {
		log.Fatal("An error has occurred 123:", err)
	}
	return bs.Bytes()
}

const (
	CLASS_IN uint16 = 1
	TYPE_A   uint16 = 1
)

func New(domainName string, recordType uint16) []byte {
	bs := make([]byte, 0)

	name := encodeDNSName(domainName)

	// rand.Seed(1)

	// id := rand.Intn(65535)
	id := 17611
	RECURSION_DESIRED := 1 << 8

	header := DNSHeader{Id: uint16(id), NumQuestions: 1, Flags: uint16(RECURSION_DESIRED)}
	question := DNSQuestion{Name: name, Type: recordType, Class: CLASS_IN}

	bs = append(bs, header.ToBytes()...)
	bs = append(bs, question.ToBytes()...)

	return bs
}
