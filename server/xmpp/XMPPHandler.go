package xmpp

import (
    "log"
    "os"
    "os/signal"
    "syscall"
	"errors"
    "fmt"
    "encoding/xml"
    "strings"
    "crypto/tls"
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
    // Read and handle the initial <stream:features> response
    buffer := make([]byte, 4096)
    n, err := conn.Conn.Read(buffer)
    if err != nil {
        return fmt.Errorf("error reading initial response: %v", err)
    }

    initialResponse := string(buffer[:n])
    log.Printf("Received initial response: %s\n", initialResponse)

    if strings.Contains(initialResponse, "<stream:features>") {
        log.Println("Stream features received, continuing with registration...")
        // You can add specific logic here to handle features if needed
    } else {
        return errors.New("expected <stream:features> but did not receive it")
    }

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

    // Wait for the actual IQ response
    n, err = conn.Conn.Read(buffer)
    if err != nil {
        return fmt.Errorf("error reading registration response: %v", err)
    }

    response := string(buffer[:n])
    log.Printf("Received registration response: %s\n", response)

    iqResponse, err := ParseIQ([]byte(response))
    if err != nil {
        return fmt.Errorf("failed to parse IQ response: %v", err)
    }

    // Handle registration error (e.g., user already exists)
    if iqResponse.IsError() {
        if iqResponse.Type == "error" {
            var iqError struct {
                Code    int    `xml:"error>code,attr"`
                Message string `xml:"error>text,omitempty"`
                Type    string `xml:"error>conflict"`
            }
            if err := xml.Unmarshal(buffer[:n], &iqError); err != nil {
                return fmt.Errorf("failed to unmarshal IQ error: %v", err)
            }

            switch iqError.Code {
            case 409:
                return fmt.Errorf("user already exists")
            default:
                return fmt.Errorf("registration failed with error code %d: %s", iqError.Code, iqError.Message)
            }
        }
    } else if iqResponse.IsResult() {
        log.Println("User created successfully")
        return nil
    }

    return errors.New("unexpected registration response")
}


// StartTLS sends the STARTTLS command to the server and upgrades the connection to TLS.
func StartTLS(conn *XMPPConnection) error {
    startTLS := "<starttls xmlns='urn:ietf:params:xml:ns:xmpp-tls'/>"
    if _, err := conn.Conn.Write([]byte(startTLS)); err != nil {
        return fmt.Errorf("failed to send STARTTLS: %v", err)
    }

    buffer := make([]byte, 4096)
    n, err := conn.Conn.Read(buffer)
    if err != nil {
        return fmt.Errorf("failed to read STARTTLS response: %v", err)
    }

    response := string(buffer[:n])
    log.Printf("Received STARTTLS response: %s\n", response)

    if strings.Contains(response, "<proceed") {
        log.Println("Proceeding with TLS handshake...")
        tlsConn := tls.Client(conn.Conn, &tls.Config{
            InsecureSkipVerify: true, // Disable verification for testing, not recommended for production
        })
        conn.Conn = tlsConn
        if err := tlsConn.Handshake(); err != nil {
            return fmt.Errorf("TLS handshake failed: %v", err)
        }
        log.Println("TLS handshake successful")
        return nil
    }

    return errors.New("failed to initiate STARTTLS")
}