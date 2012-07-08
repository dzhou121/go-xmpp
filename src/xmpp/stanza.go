package xmpp

import "encoding/xml"

// XMPP <iq/> stanza.
type Iq struct {
	XMLName xml.Name `xml:"iq"`
	Id string `xml:"id,attr"`
	Type string `xml:"type,attr"`
	Payload string `xml:",innerxml"`
}

// Encode the value to an XML string and set as the payload. See xml.Marshal
// for how the value is encoded.
func (iq *Iq) PayloadEncode(v interface{}) error {
	bytes, err := xml.Marshal(v)
	if err != nil {
		return err
	}
	iq.Payload = string(bytes)
	return nil
}

// Decode the payload (an XML string) into the given value. See xml.Unmarshal
// for how the value is decoded.
func (iq *Iq) PayloadDecode(v interface{}) error {
	return xml.Unmarshal([]byte(iq.Payload), v)
}

// XMPP <message/> stanza.
type Message struct {
	XMLName xml.Name `xml:"message"`
}

// XMPP <presence/> stanza.
type Presence struct {
	XMLName xml.Name `xml:"presence"`
}
