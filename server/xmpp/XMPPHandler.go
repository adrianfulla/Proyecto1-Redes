package xmpp

import (
    "log"
    "os"
    "os/signal"
    "syscall"
	"errors"
    "fmt"
)

type XMPPHandler struct {
    Server   string
    Username string
    Password string
    Conn     *XMPPConnection
}

// NewXMPPHandler creates and initializes an XMPPHandler
func NewXMPPHandler(server, username, password string) (*XMPPHandler, error) {
    handler := &XMPPHandler{
        Server:   server,
        Username: username,
        Password: password,
    }

    conn, err := NewXMPPConnection(server, false)
    if err != nil {
        return nil, err
    }
    handler.Conn = conn

    if err := handler.startStream(); err != nil {
        return nil, err
    }

    if err := Authenticate(conn, username, password); err != nil {
        return nil, err
    }

    return handler, nil
}

// startStream initializes the XMPP stream for the handler
func (xh *XMPPHandler) startStream() error {
    return xh.Conn.StartStream(xh.Username)
}

// SendMessage sends a chat message to a specified recipient
func (xh *XMPPHandler) SendMessage(to, body string) error {
    message := NewMessage(to, "chat", body)
    return xh.sendStanza(message)
}

// SendPresence sends a presence stanza (e.g., available, unavailable)
func (xh *XMPPHandler) SendPresence(presenceType, status string) error {
    presence := NewPresence("", presenceType, status, "", 0)
    return xh.sendStanza(presence)
}

// SendIQ sends an IQ stanza
func (xh *XMPPHandler) SendIQ(iqType, id string, query interface{}) error {
    iq := NewIQ(iqType, id)
    iq.SetQuery(query)
    return xh.sendStanza(iq)
}

// sendStanza converts a stanza to XML and sends it over the connection
func (xh *XMPPHandler) sendStanza(stanza Stanza) error {
    xml, err := stanza.ToXML()
    if err != nil {
        return err
    }
    _, err = xh.Conn.Conn.Write([]byte(xml))
    return err
}

// HandleIncomingStanzas listens and processes incoming stanzas
func (xh *XMPPHandler) HandleIncomingStanzas() {
    go func() {
        for {
            buffer := make([]byte, 4096)
            n, err := xh.Conn.Conn.Read(buffer)
            if err != nil {
                log.Printf("Error reading from connection: %v", err)
                return
            }

            stanza := string(buffer[:n])
            fmt.Printf("Received stanza: %s\n", stanza)

            if iq, err := ParseIQ([]byte(stanza)); err == nil {
                response, err := HandleIQ(iq)
                if err != nil {
                    log.Printf("Error handling IQ: %v", err)
                } else if response != nil {
                    xh.sendStanza(response)
                }
            } else if msg, err := ParseMessage([]byte(stanza)); err == nil {
                fmt.Printf("Received message from %s: %s\n", msg.From, msg.Body)
            } else if presence, err := ParsePresence([]byte(stanza)); err == nil {
                fmt.Printf("Received presence from %s: %s\n", presence.From, presence.Status)
            } else {
                log.Printf("Unknown stanza received: %s", stanza)
            }
        }
    }()
}

// WaitForShutdown handles graceful shutdowns on receiving a signal
func (xh *XMPPHandler) WaitForShutdown() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    <-c
    fmt.Println("Shutting down...")

    xh.SendPresence("unavailable", "")

    xh.Conn.CloseStream()
    xh.Conn.Close()
}


// CreateUser attempts to create a new user on the XMPP server.
func CreateUser(conn *XMPPConnection, username, password string) error {
    // Create a registration request IQ stanza
    iqID := "register1" // This should be unique for each request in a real application
    register := RegisterRequest{
        Username: username,
        Password: password,
    }
    
    iq := NewIQ("set", iqID)
    iq.SetQuery(register)

    // Send the registration request
    if err := sendStanza(conn, iq); err != nil {
        return fmt.Errorf("failed to send registration request: %v", err)
    }

    // Wait for the response
    buffer := make([]byte, 4096)
    n, err := conn.Conn.Read(buffer)
    if err != nil {
        return fmt.Errorf("error reading registration response: %v", err)
    }

    // Parse the response
    response := string(buffer[:n])
    fmt.Printf("Received registration response: %s\n", response)

    if iqResponse, err := ParseIQ([]byte(response)); err == nil {
        if iqResponse.IsError() {
            return errors.New("registration failed")
        } else if iqResponse.IsResult() {
            fmt.Println("Registration successful")
            return nil
        }
    }

    return errors.New("unexpected registration response")
}