package protocol

import "io"

type IPacket interface {
	Pack(writer io.Writer) error
	Unpack(reader io.Reader) error
}

type IPacketHandler interface {
	Packet(packet IPacket)
	Error(err error)
}
