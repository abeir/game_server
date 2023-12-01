package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

var mark = [4]byte{'G', 'S', 'A', '0'}

func Mark() [4]byte {
	return mark
}

type MarkPacket struct {
	Mark   [4]byte
	Length int32
	Body   []byte
}

func (m *MarkPacket) Pack(writer io.Writer) error {
	var err error
	if err = binary.Write(writer, binary.BigEndian, &m.Mark); err != nil {
		return err
	}
	if err = binary.Write(writer, binary.BigEndian, &m.Length); err != nil {
		return err
	}
	err = binary.Write(writer, binary.BigEndian, &m.Body)
	return err
}

func (m *MarkPacket) Unpack(reader io.Reader) error {
	var err error
	if err = binary.Read(reader, binary.BigEndian, &m.Mark); err != nil {
		return err
	}
	if err = binary.Read(reader, binary.BigEndian, &m.Length); err != nil {
		return err
	}
	m.Body = make([]byte, m.Length)
	err = binary.Read(reader, binary.BigEndian, &m.Body)
	return err
}

func CreateMarkPacket(mark [4]byte, body []byte) MarkPacket {
	pack := MarkPacket{
		Mark:   mark,
		Length: int32(len(body)),
		Body:   body,
	}
	return pack
}

func MarkPacketSplit(mark [4]byte, reader io.Reader, handle IPacketHandler) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		if !atEOF && bytes.Compare(data[:4], mark[:]) == 0 {
			if len(data) > 8 {
				length := int32(0)
				err = binary.Read(bytes.NewReader(data[4:8]), binary.BigEndian, &length)

				fullLength := length + 8
				if fullLength <= int32(len(data)) {
					return int(fullLength), data[:fullLength], nil
				}
			}
		}
		return
	})
	for scanner.Scan() {
		scannedPack := new(MarkPacket)
		err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			handle.Error(err)
			continue
		}
		handle.Packet(scannedPack)
	}
	return scanner.Err()
}
