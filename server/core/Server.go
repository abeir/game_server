package core

import (
	"game_server/server/common/conf"
	"game_server/server/common/log"
	"game_server/server/common/utils"
	"io"
	"net"
)

type ISender interface {
	Send(writer io.Writer, data []byte)
	Broadcast(data []byte)
}

type Server struct {
	logger     log.ILogger
	listener   *net.TCPListener
	dispatcher IDispatcher

	setting conf.Server
	started bool
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SetDispatcher(dispatcher IDispatcher) {
	s.dispatcher = dispatcher
}

func (s *Server) SetLogger(logger log.ILogger) {
	s.logger = logger
}

func (s *Server) Setting(setting conf.Server) {
	s.setting = setting
}

func (s *Server) Start() error {
	if s.started {
		return utils.InitializeError("server started")
	}
	var err error
	s.listener, err = net.ListenTCP("tcp", &net.TCPAddr{
		IP:   []byte(s.setting.Listen),
		Port: s.setting.Port,
	})
	if err != nil {
		return err
	}
	s.started = true

	for {
		conn, err := s.listener.AcceptTCP()
		if err != nil {
			s.logger.Error(err)
			continue
		}
		s.process(conn)
	}
}

func (s *Server) Shutdown() {
	go func() {
		if err := s.listener.Close(); err != nil {
			s.logger.Error(err)
		} else {
			s.dispatcher.Dispose()
		}
	}()
}

func (s *Server) process(conn *net.TCPConn) {
	s.dispatcher.AddConn(conn)

	err := s.dispatcher.Dispatch(conn)
	if err != nil {
		s.logger.Error(err)
	}
}
