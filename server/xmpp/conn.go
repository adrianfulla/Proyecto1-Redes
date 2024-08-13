package xmpp

import (
	"crypto/tls"
	"net"
)

type XMPPConnection struct {
    Conn net.Conn
}

func (c *XMPPConnection) Write(data []byte) (int, error) {
	return c.Conn.Write(data)
}

func (c *XMPPConnection) Read(buffer []byte) (int, error) {
	return c.Conn.Read(buffer)
}

func NewXMPPConnection(server string, useTLS bool) (*XMPPConnection, error) {
	var conn net.Conn
	var err error

	if useTLS {
		conn, err = tls.Dial("tcp", server, &tls.Config{InsecureSkipVerify: true})
	} else {
		conn, err = net.Dial("tcp", server)
	}

	if err != nil {
		return nil, err
	}

	return &XMPPConnection{Conn: conn}, nil
}

func (xc *XMPPConnection) Close() error {
    return xc.Conn.Close()
}
