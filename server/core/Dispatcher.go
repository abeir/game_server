package core

import (
	"game_server/server/common/log"
	"game_server/server/common/utils"
	"game_server/server/core/protocol"
	"net"
	"sync"
)

type IDispatcher interface {
	protocol.IPacketHandler

	SetLogger(logger log.ILogger)

	AddConn(conn *net.TCPConn)
	Dispose()

	Dispatch(conn *net.TCPConn) error
}

type Dispatcher struct {
	logger      log.ILogger
	connections sync.Map
	gen         utils.Generator
}

func (d *Dispatcher) SetLogger(logger log.ILogger) {
	d.logger = logger
}

func (d *Dispatcher) AddConn(conn *net.TCPConn) {
	id := d.gen.Next()

	d.connections.Store(id, conn)
}

func (d *Dispatcher) Dispose() {
}

func (d *Dispatcher) Dispatch(conn *net.TCPConn) error {
	return protocol.MarkPacketSplit(protocol.Mark(), conn, d)
}

func (d *Dispatcher) Packet(packet protocol.IPacket) {
	//pack := packet.(*protocol.MarkPacket)
}

func (d *Dispatcher) Error(err error) {
	d.logger.Error(err)
}
