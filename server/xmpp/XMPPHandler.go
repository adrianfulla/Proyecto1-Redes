package xmpp

import (
    "log"
    "fmt"
    "bufio"
)

type XMPPHandler struct {
    Conn     *XMPPConnection
    Server   string
    Username string
    Password string
}

func NewXMPPHandler(domain, port, username, password string) (*XMPPHandler, error) {
    handler := &XMPPHandler{
        Server:   domain +":"+port,
        Username: username,
        Password: password,
    }

    conn, err := NewXMPPConnection(domain, port, false)
    if err != nil {
        return nil, err
    }
    handler.Conn = conn

    if err := handler.Conn.StartStream(""); err != nil {
        return nil, err
    }

    // If the user does not exist, create it
    if err := CreateUser(handler.Conn, username, password); err != nil {
        log.Println("User creation failed or user already exists, proceeding with login...")
    }

    // Authenticate
    if err := Authenticate(handler.Conn, username, password); err != nil {
        return nil, err
    }

    // Bind Resource
    if err := BindResource(handler.Conn); err != nil {
        return nil, err
    }

    return handler, nil
}


// SendPresence sends a presence stanza to update the user's availability status.
func (h *XMPPHandler) SendPresence(presenceType, status string) error {
    presence := fmt.Sprintf(
        `<presence><show>%s</show><status>%s</status></presence>`,
        presenceType, status,
    )
    _, err := h.Conn.Conn.Write([]byte(presence))
    if err != nil {
        log.Printf("Failed to send presence: %v", err)
        return err
    }
    log.Println("Presence sent successfully")
    return nil
}

// SendMessage sends a message stanza to the specified recipient.
func (h *XMPPHandler) SendMessage(to, message string) error {
    msg := fmt.Sprintf(
        `<message to='%s' type='chat'><body>%s</body></message>`,
        to, message,
    )
    _, err := h.Conn.Conn.Write([]byte(msg))
    if err != nil {
        log.Printf("Failed to send message: %v", err)
        return err
    }
    log.Println("Message sent successfully")
    return nil
}

// HandleIncomingStanzas listens for incoming stanzas and processes them.
func (h *XMPPHandler) HandleIncomingStanzas() {
    reader := bufio.NewReader(h.Conn.Conn)
    for {
        stanza, err := reader.ReadString('>')
        if err != nil {
            log.Printf("Failed to read stanza: %v", err)
            return
        }
        log.Printf("Received stanza: %s", stanza)
        // Here you would parse and handle the stanza based on its type
    }
}

// WaitForShutdown keeps the connection alive until shutdown is requested.
func (h *XMPPHandler) WaitForShutdown() {
    log.Println("Waiting for shutdown signal...")
    // Implementation could wait on a signal or just block until interrupted
    select {}
}