package core

import (
	"net"
)

// TCPConnHandler handles TCP connections comming from TUN.
type TCPConnHandler interface {
	// Handle handles the conn for target.
	Handle(conn net.Conn, target *net.TCPAddr) error
}

// UDPConnHandler handles UDP connections comming from TUN.
type UDPConnHandler interface {
	// Connect connects the proxy server. Note that target can be nil.
	Connect(conn UDPConn, target *net.UDPAddr) error

	// ReceiveTo will be called when data arrives from TUN.
	ReceiveTo(conn UDPConn, data []byte, addr *net.UDPAddr) error
}

// Interface implemented by TUNHandler.
type UDPHandler interface {
	HandleUdp(dstAddr net.IP, dstPort uint16,
			localAddr net.IP, localPort uint16,
			data []byte)
}

var tcpConnHandler TCPConnHandler
var udpConnHandler UDPConnHandler
var rawUdpConnHandler UDPHandler

func RegisterTCPConnHandler(h TCPConnHandler) {
	tcpConnHandler = h
}

func RegisterUDPConnHandler(h UDPConnHandler) {
	udpConnHandler = h
}

func RegisterRawUDPHandler(h UDPHandler) {
	rawUdpConnHandler = h
}
