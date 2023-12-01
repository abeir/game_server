package testcase

import (
	s "game_server/server"
	"testing"
)

func TestServer(t *testing.T) {
	println("hello test")

	ret := s.Server()

	if ret != "Server" {
		t.Errorf("Server() return expect: %s, but: %s", "Server", ret)
	}
}
