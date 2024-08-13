package xmpp

import (
    "encoding/xml"
    "errors"
    "fmt"
)

// AuthRequest represents the structure of an authentication request in XMPP.
type AuthRequest struct {
    XMLName  xml.Name `xml:"jabber:iq:auth query"`
    Username string   `xml:"username"`
    Password string   `xml:"password"`
    Resource string   `xml:"resource,omitempty"`
}

// Authenticate performs simple plain-text authentication with the XMPP server.
func Authenticate(conn *XMPPConnection, username, password string) error {
    // Create an authentication request IQ stanza
    iqID := "auth1" // This should be unique for each request in a real application
    auth := AuthRequest{
        Username: username,
        Password: password,
    }
    
    iq := NewIQ("set", iqID)
    iq.SetQuery(auth)

    // Send the authentication request
    if err := sendStanza(conn, iq); err != nil {
        return fmt.Errorf("failed to send authentication request: %v", err)
    }

    // Wait for the response
    buffer := make([]byte, 4096)
    n, err := conn.Conn.Read(buffer)
    if err != nil {
        return fmt.Errorf("error reading authentication response: %v", err)
    }

    // Parse the response
    response := string(buffer[:n])
    fmt.Printf("Received authentication response: %s\n", response)

    if iqResponse, err := ParseIQ([]byte(response)); err == nil {
        if iqResponse.IsError() {
            return errors.New("authentication failed")
        } else if iqResponse.IsResult() {
            fmt.Println("Authentication successful")
            return nil
        }
    }

    return errors.New("unexpected authentication response")
}

// sendStanza is a helper function to send an XMPP stanza over the connection
func sendStanza(conn *XMPPConnection, stanza Stanza) error {
    xml, err := stanza.ToXML()
    if err != nil {
        return err
    }
    _, err = conn.Conn.Write([]byte(xml))
    return err
}



type RegisterRequest struct {
    XMLName  xml.Name `xml:"jabber:iq:register query"`
    Username string   `xml:"username"`
    Password string   `xml:"password"`
}