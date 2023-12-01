package testcase

import (
	"bytes"
	"game_server/server/core/protocol"
	"testing"
)

const testMarkPaketBody = "this is a test body"

func createByteBuffer(mark [4]byte) (*bytes.Buffer, error) {
	pack := protocol.CreateMarkPacket(mark, []byte(testMarkPaketBody))

	buf := new(bytes.Buffer)
	var err error
	for i := 0; i < 6; i++ {
		if err = pack.Pack(buf); err != nil {
			return nil, err
		}
	}
	return buf, nil
}

type testPacketHandlerImpl struct {
	t *testing.T
}

func (t *testPacketHandlerImpl) Packet(packet protocol.IPacket) {
	pack := packet.(*protocol.MarkPacket)
	if string(pack.Body) != testMarkPaketBody {
		t.t.Errorf("body invalid, expect: %s, but: %s", testMarkPaketBody, string(pack.Body))
	}
}

func (t *testPacketHandlerImpl) Error(err error) {
	if err != nil {
		t.t.Error(err)
	}
}

func TestMarkPacketSplit(t *testing.T) {
	mark := [4]byte{'T', '_', '0', '1'}

	buf, err := createByteBuffer(mark)
	if err != nil {
		t.Error(err)
		return
	}
	handler := &testPacketHandlerImpl{t}
	err = protocol.MarkPacketSplit(mark, buf, handler)
	if err != nil {
		t.Error(err)
		return
	}
}
