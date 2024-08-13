package xmpp

import (
    "encoding/xml"
    "errors"
)

// IQ represents an XMPP IQ stanza.
type IQ struct {
    XMLName xml.Name `xml:"iq"`
    From    string   `xml:"from,attr,omitempty"`
    To      string   `xml:"to,attr,omitempty"`
    Type    string   `xml:"type,attr"` // "get", "set", "result", "error"
    ID      string   `xml:"id,attr"`
    Query   interface{} `xml:",innerxml"` // Contains the query or command payload
}

// NewIQ creates a new IQ stanza with the specified type and ID.
func NewIQ(iqType, id string) *IQ {
    return &IQ{
        Type: iqType,
        ID:   id,
    }
}

// ToXML converts the IQ struct to an XML string.
func (iq *IQ) ToXML() (string, error) {
    output, err := xml.Marshal(iq)
    if err != nil {
        return "", err
    }
    return string(output), nil
}

// ParseIQ parses an XML string into an IQ struct.
func ParseIQ(data []byte) (*IQ, error) {
    var iq IQ
    if err := xml.Unmarshal(data, &iq); err != nil {
        return nil, err
    }
    return &iq, nil
}

// SetQuery sets the query or command payload for the IQ stanza.
func (iq *IQ) SetQuery(query interface{}) {
    iq.Query = query
}

// Additional utility functions

// IsGet checks if the IQ type is "get".
func (iq *IQ) IsGet() bool {
    return iq.Type == "get"
}

// IsSet checks if the IQ type is "set".
func (iq *IQ) IsSet() bool {
    return iq.Type == "set"
}

// IsResult checks if the IQ type is "result".
func (iq *IQ) IsResult() bool {
    return iq.Type == "result"
}

// IsError checks if the IQ type is "error".
func (iq *IQ) IsError() bool {
    return iq.Type == "error"
}

// HandleIQ processes incoming IQ stanzas based on their type and returns a response.
func HandleIQ(iq *IQ) (*IQ, error) {
    switch {
    case iq.IsGet():
        // Handle the "get" IQ stanza.
        // Here you could route to specific handlers based on the query type.
        // For example, you might have a service discovery handler, a vCard handler, etc.
        return processIQGet(iq)
    case iq.IsSet():
        // Handle the "set" IQ stanza.
        return processIQSet(iq)
    case iq.IsResult():
        // Handle the "result" IQ stanza.
        return nil, nil // Typically, results are responses to previous "get" or "set" requests
    case iq.IsError():
        // Handle the "error" IQ stanza.
        return nil, errors.New("received error IQ stanza")
    default:
        return nil, errors.New("unknown IQ type")
    }
}

// processIQGet processes a "get" IQ request.
func processIQGet(iq *IQ) (*IQ, error) {
    // Example: Check if the query is a service discovery request
    // You'd typically decode the query field into a more specific struct here
    return nil, nil
}

// processIQSet processes a "set" IQ request.
func processIQSet(iq *IQ) (*IQ, error) {
    // Example: Handle setting a value, like setting a vCard or roster item
    return nil, nil
}
