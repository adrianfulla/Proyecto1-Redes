package xmpp

import (
	"crypto/tls"
	"net"
	"time"
	"fmt"
)

type XMPPConnection struct {
    Conn net.Conn
}

func NewXMPPConnection(address string, useTLS bool) (*XMPPConnection, error) {
    var conn net.Conn
    var err error

	dialer := &net.Dialer{
        Timeout: 5 * time.Second,
    }

    if useTLS {
		// TLS config with InsecureSkipVerify to bypass certificate verification
        tlsConfig := &tls.Config{
            InsecureSkipVerify: true, // Disable certificate verification
            ServerName:         "alumchat.lol",
        }

        // Directly establish a TLS connection, effectively skipping STARTTLS
        conn, err = tls.DialWithDialer(dialer, "tcp", address, tlsConfig)
        if err != nil {
            return nil, fmt.Errorf("TLS connection failed: %w", err)
        }
    } else {
        conn, err = net.Dial("tcp", address)
    }

    if err != nil {
        return nil, err
    }

    return &XMPPConnection{Conn: conn}, nil
}

func (xc *XMPPConnection) Close() error {
    return xc.Conn.Close()
}
