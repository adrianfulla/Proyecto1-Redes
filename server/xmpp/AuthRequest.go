package xmpp

import (
    "encoding/xml"
    "encoding/base64"
    "errors"
    "fmt"
    "strings"
)

type AuthRequest struct {
    XMLName   xml.Name `xml:"auth"`
    Mechanism string   `xml:"mechanism,attr"`
    Text      string   `xml:",chardata"`
}

// ToXML converts the AuthRequest to its XML representation.
func (a *AuthRequest) ToXML() (string, error) {
    output, err := xml.Marshal(a)
    if err != nil {
        return "", fmt.Errorf("failed to marshal AuthRequest: %v", err)
    }
    return string(output), nil
}

// Authenticate performs SASL authentication with the XMPP server.
// func Authenticate(conn *XMPPConnection, username, password string) error {
//     // Perform STARTTLS
//     if err := StartTLS(conn); err != nil {
//         return fmt.Errorf("STARTTLS failed: %v", err)
//     }

//     // After STARTTLS, reinitiate the stream and expect <stream:features> again
//     if err := conn.StartStream(""); err != nil {
//         return fmt.Errorf("failed to start stream after STARTTLS: %v", err)
//     }

//     // Read and handle the initial <stream:features> response
//     buffer := make([]byte, 4096)
//     n, err := conn.Conn.Read(buffer)
//     if err != nil {
//         return fmt.Errorf("error reading initial response after STARTTLS: %v", err)
//     }

//     initialResponse := string(buffer[:n])
//     log.Printf("Received initial response after STARTTLS: %s\n", initialResponse)

//     if strings.Contains(initialResponse, "<stream:features>") {
//         log.Println("Stream features received, proceeding with authentication...")
//     } else {
//         return errors.New("expected <stream:features> but did not receive it after STARTTLS")
//     }

//     // Perform SASL PLAIN authentication
//     authText := "\x00" + username + "\x00" + password
//     authBase64 := base64.StdEncoding.EncodeToString([]byte(authText))
//     auth := AuthRequest{
//         Mechanism: "PLAIN",
//         Text:      authBase64,
//     }

//     if err := sendStanza(conn, &auth); err != nil {
//         return fmt.Errorf("failed to send authentication request: %v", err)
//     }

//     // Wait for the response
//     n, err = conn.Conn.Read(buffer)
//     if err != nil {
//         return fmt.Errorf("error reading authentication response: %v", err)
//     }

//     response := string(buffer[:n])
//     log.Printf("Received authentication response: %s\n", response)

//     // Parse the response and check for success
//     if strings.Contains(response, "<success") {
//         log.Println("Authentication successful")
//         return nil
//     }

//     return errors.New("authentication failed or unexpected response")
// }

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

// Authenticate performs SASL PLAIN authentication with the XMPP server.
func Authenticate(conn *XMPPConnection, username, password string) error {
	authStr := "\x00" + username + "\x00" + password
	authBase64 := base64.StdEncoding.EncodeToString([]byte(authStr))

	authStanza := fmt.Sprintf(`<auth xmlns="urn:ietf:params:xml:ns:xmpp-sasl" mechanism="PLAIN">%s</auth>`, authBase64)

	// Send the authentication stanza
	if _, err := conn.Write([]byte(authStanza)); err != nil {
		return fmt.Errorf("failed to send authentication request: %v", err)
	}

	// Read the server response
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("error reading authentication response: %v", err)
	}

	response := string(buffer[:n])
	fmt.Printf("Received authentication response: %s\n", response)

	// Check if authentication was successful
	if strings.Contains(response, "<success") {
		fmt.Println("Authentication successful")
		return nil
	}

	// Handle authentication failure
	if strings.Contains(response, "<failure") {
		return errors.New("authentication failed")
	}

	return errors.New("unexpected authentication response")
}

